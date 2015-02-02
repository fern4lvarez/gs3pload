package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/docopt/docopt-go"
	"github.com/mitchellh/go-homedir"
)

var (
	HOME, _           = homedir.Dir()
	ENVIRONMENTS_FILE = filepath.Join(HOME, ".gs3pload", "envs.json")
	VERSION           = "0.0.1"
)

func setBucket(name string, environmentType string, environmentName string) string {
	var path, root string

	pathSplit := strings.SplitN(name, "/", 2)
	name = pathSplit[0]
	switch name {
	case "packages":
		root = fmt.Sprintf("%s.%s", name, environmentName)
	case "certs", "images", "stacks":
		root = fmt.Sprintf("%s.%s", environmentName, name)
	default:
		root = name
	}

	if len(pathSplit) > 1 {
		path = fmt.Sprintf("%s/%s", root, pathSplit[1])
	} else {
		path = root
	}

	return fmt.Sprintf("%s://%s/", environmentType, path)
}

// Execute executes a regular command splitted in strings
func Execute(command []string) error {
	cmd := exec.Command(command[0], command[1:]...)

	cmd.Stdout = os.Stdout
	return cmd.Run()
}

// Copy files to a given bucket and environment
func Copy(config string, bucket string, files []string, recursive bool, environment Environment) error {
	command := []string{"gsutil", "cp"}
	if recursive {
		command = append(command, "-R")
	}
	command = append(command, files...)
	command = append(command, bucket)

	os.Setenv("BOTO_CONFIG", config)
	return Execute(command)
}

// Public set public-read permissions for the given files
func Public(config string, bucket string, files []string) error {
	command := []string{"gsutil", "acl", "set", "public-read"}
	for _, file := range files {
		filePath := fmt.Sprintf("%s%s", bucket, file)
		command = append(command, filePath)
	}
	os.Setenv("BOTO_CONFIG", config)
	return Execute(command)
}

// DaisyChain copy an object from one path to another, where these
// can belong to different buckets or environments
func DaisyChain(config string, originPath, destPath string, recursive bool) error {
	command := []string{"gsutil", "cp", "-D", "-p"}
	if recursive {
		command = append(command, "-R")
	}
	command = append(command, originPath)
	command = append(command, destPath)
	os.Setenv("BOTO_CONFIG", config)
	return Execute(command)
}

// Backup given files within the same bucket with a .bak extension
func Backup(config string, files []string, bucket string, recursive bool) {
	fmt.Printf("---> Creating backups on %s...\n", bucket)
	for _, file := range files {
		filePath := fmt.Sprintf("%s%s", bucket, file)
		backupPath := fmt.Sprintf("%s%s%s", bucket, file, ".bak")
		if err := DaisyChain(config, filePath, backupPath, recursive); err != nil {
			fmt.Printf("Skipping backup of %s (Did it exist before?). \n", filePath)
			continue
		}
	}
}

// Push given files to multiple environments
func Push(environments Environments, bucketName string, files []string, recursive, public, backup bool) error {
	for _, environment := range environments {
		fmt.Printf("---> Pushing to %s environment...\n", environment.Name)
		bucket := setBucket(bucketName, environment.Type, environment.Name)
		config := fmt.Sprintf("%s.boto", filepath.Join(HOME, ".gs3pload", environment.Name))

		if backup {
			Backup(config, files, bucket, recursive)
		}

		if err := Copy(config, bucket, files, recursive, environment); err != nil {
			fmt.Printf("Push failed on %s. %s\n", environment.Name, err)
			continue
		}

		if public {
			if err := Public(config, bucket, files); err != nil {
				fmt.Printf("Set as public failed on %s. %s\n", environment.Name, err)
				continue
			}
		}
	}
	return nil
}

func main() {
	usage := `gs3pload. Upload files to multiple S3 or Google Storage
buckets at once.

Bucket names "packages", "stacks", "certs" and "images" are reserved and resolved
by environment domain.

Usage:
  gs3pload push <bucket> <name>... [-r | --recursive] [-p | --public] [-b | --backup]
  gs3pload -h | --help
  gs3pload -v | --version

Options:
  -h --help        Show help.
  -p --public      Set files as public.
  -r --recursive   Do a recursive copy.
  -b --backup      Create backup of pushed files if they exist.
  -v --version     Show version.
`

	arguments, _ := docopt.Parse(usage, nil, true, fmt.Sprintf("gs3pload %s", VERSION), false)

	push := arguments["push"].(bool)
	bucketName := arguments["<bucket>"].(string)
	fileNames := arguments["<name>"].([]string)
	recursive := arguments["--recursive"].(bool)
	public := arguments["--public"].(bool)
	backup := arguments["--backup"].(bool)

	environments := Environments{}
	err := environments.Fetch()
	if err != nil {
		fmt.Println(err)
		return
	}

	if push {
		err = Push(environments, bucketName, fileNames, recursive, public, backup)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
