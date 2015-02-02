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
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// Copy files to a given bucket and environment
func Copy(config string, bucket string, files []string, environment Environment) error {
	command := []string{"gsutil", "cp"}
	command = append(command, files...)
	command = append(command, bucket)

	os.Setenv("BOTO_CONFIG", config)
	return Execute(command)
}

// Public set public-read permissions for the given files
func Public(config string, bucket string, files []string, environment Environment) error {
	command := []string{"gsutil", "acl", "set", "public-read"}
	for _, file := range files {
		filePath := fmt.Sprintf("%s%s", bucket, file)
		command = append(command, filePath)
	}
	os.Setenv("BOTO_CONFIG", config)
	return Execute(command)
}

// Push given files to multiple environments
func Push(environments Environments, bucketName string, files []string, public bool) error {
	for _, environment := range environments {
		fmt.Printf("---> Pushing to %s environment...\n", environment.Name)
		bucket := setBucket(bucketName, environment.Type, environment.Name)
		config := fmt.Sprintf("%s.boto", filepath.Join(HOME, ".gs3pload", environment.Name))

		err := Copy(config, bucket, files, environment)
		if err != nil {
			fmt.Printf("Push failed on %s. %s\n", environment.Name, err)
			continue
		}

		if public {
			err = Public(config, bucket, files, environment)
			if err != nil {
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
  gs3pload push <bucket> <name>... [-p | --public]
  gs3pload -h | --help
  gs3pload -v | --version

Options:
  -h --help        Show help.
  -p --public      Set files as public.
  -v --version     Show version.
`

	arguments, _ := docopt.Parse(usage, nil, true, fmt.Sprintf("gs3pload %s", VERSION), false)

	push := arguments["push"].(bool)
	bucketName := arguments["<bucket>"].(string)
	fileNames := arguments["<name>"].([]string)
	public := arguments["--public"].(bool)

	environments := Environments{}
	err := environments.Fetch()
	if err != nil {
		fmt.Println(err)
		return
	}

	if push {
		err = Push(environments, bucketName, fileNames, public)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
