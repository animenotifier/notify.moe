package main

import (
	"encoding/json"
	"io/ioutil"
)

// APIKeys ...
type APIKeys struct {
	Google struct {
		ID     string `json:"id"`
		Secret string `json:"secret"`
	} `json:"google"`

	Facebook struct {
		ID     string `json:"id"`
		Secret string `json:"secret"`
	} `json:"facebook"`
}

var apiKeys APIKeys

func init() {
	data, _ := ioutil.ReadFile("security/api-keys.json")
	err := json.Unmarshal(data, &apiKeys)

	if err != nil {
		panic(err)
	}
}
