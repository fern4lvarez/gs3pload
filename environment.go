package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Environment struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type Environments []Environment

func (environments *Environments) Fetch(envs interface{}) error {
	environmentsFile := ENVIRONMENTS_FILE
	if envs != nil {
		environmentsFile = envs.(string)
	}

	b, err := ioutil.ReadFile(environmentsFile)
	if err != nil {
		return err
	}

	if err = json.Unmarshal(b, &environments); err != nil {
		return err
	}

	return nil
}

func (environment *Environment) Prepare(file string) error {
	if environment.getBackend() == "gsutil" {
		if err := os.Setenv("BOTO_CONFIG", file); err != nil {
			return err
		}
	}

	return nil
}

func (environment *Environment) getBackend() string {
	if environment.Type == "s3" || environment.Type == "gs" {
		return "gsutil"
	}

	return "swift"
}
