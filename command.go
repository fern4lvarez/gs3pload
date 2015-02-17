package main

import (
	"fmt"
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
}
