package main

import (
	"encoding/json"
	"io/ioutil"
)

type Environment struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type Environments []Environment

func (environments *Environments) Fetch(file interface{}) error {
	environmentsFile := ENVIRONMENTS_FILE
	if file != nil {
		environmentsFile = file.(string)
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
