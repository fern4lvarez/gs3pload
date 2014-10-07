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

func (environments *Environments) Fetch() error {
	b, err := ioutil.ReadFile(ENVIRONMENTS_FILE)
	if err != nil {
		return err
	}

	if err = json.Unmarshal(b, &environments); err != nil {
		return err
	}

	return nil
}
