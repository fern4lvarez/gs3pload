package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
)

type Command struct {
	Base []string
	Type string
}

func NewCommand(t string) *Command {
	return &Command{
		Base: []string{},
		Type: t}
}

// Execute executes a command
func (command *Command) Execute() error {
	if len(command.Base) < 2 {
		return errors.New("command is too short")
	}

	cmd := exec.Command(command.Base[0], command.Base[1:]...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func (command *Command) Copy(bucket string, files []string, recursive bool) {
	if command.Type == "gsutil" {
		command.Base = append(command.Base, []string{
			"gsutil",
			"cp"}...)
		if recursive {
			command.Base = append(command.Base, "-R")
		}

		command.Base = append(command.Base, files...)
		command.Base = append(command.Base, bucket)
	}

	if command.Type == "swift" {
		command.Base = append(command.Base, []string{
			"swift",
			"upload"}...)
		command.Base = append(command.Base, bucket)
		command.Base = append(command.Base, files...)
	}
}

func (command *Command) Public(bucket string, files []string) {
	if command.Type == "gsutil" {
		command.Base = append(command.Base, []string{
			"gsutil",
			"acl",
			"set",
			"public-read"}...)
		for _, file := range files {
			filePath := fmt.Sprintf("%s%s", bucket, file)
			command.Base = append(command.Base, filePath)
		}
	}

	if command.Type == "swift" {
		command.Base = append(command.Base, []string{
			"echo",
			"-p flag not supported for swift platforms. Skipping."}...)
	}
}

func (command *Command) DaisyChain(originPath, destPath string, recursive bool) {
	if command.Type == "gsutil" {
		command.Base = append(command.Base, []string{
			"gsutil",
			"cp",
			"-D",
			"-p"}...)
		if recursive {
			command.Base = append(command.Base, "-R")
		}

		command.Base = append(command.Base, originPath)
		command.Base = append(command.Base, destPath)
	}

	if command.Type == "swift" {
		command.Base = append(command.Base, []string{
			"echo",
			"-b flag not supported for swift platforms. Skipping."}...)
	}
}
