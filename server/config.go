package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

/*
Config is the software global config
*/
type Config struct {
	Reads  []string `json:"reads"`
	Writes []string `json:"writes"`
}

const (
	configPath = "/etc/treasury/config.json"
)

func getConfig() Config {

	defaultConfig := Config{[]string{}, []string{}}

	content, err := ioutil.ReadFile(configPath)
	if err != nil {
		return writeBack(defaultConfig)
	}

	var config Config
	err = json.Unmarshal(content, &config)
	if err != nil {
		fmt.Println("malformed config.")
		return defaultConfig
	}

	return writeBack(config)
}

func writeBack(config Config) Config {
	data, err := json.MarshalIndent(config, "", " ")
	if err != nil {
		return config
	}

	ioutil.WriteFile(configPath, data, os.ModePerm)
	return config
}
