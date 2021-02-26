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

var (
	configPaths = []string{
		"/etc/treasury/config.json",
		"config.json", // Fallback local config file
	}
)

func readFiles(filePaths []string) ([]byte, error) {
	var err error = nil
	content := []byte{}

	for _, filePath := range filePaths {
		content, err = ioutil.ReadFile(filePath)
		if err == nil {
			return content, nil
		}
	}

	return nil, err
}

func getConfig() Config {

	defaultConfig := Config{[]string{}, []string{}}

	content, err := readFiles(configPaths)
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

	ioutil.WriteFile(configPaths[0], data, os.ModePerm)
	return config
}
