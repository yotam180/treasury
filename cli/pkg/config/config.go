package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
)

func getCofigPath() string {
	appdataPath := "/etc/"

	if runtime.GOOS == "windows" {
		appdataPath = os.Getenv("APPDATA")
	}

	return filepath.Join(appdataPath, "treasury-cli", "config.json")
}

/*
Configuration is used for configuration
*/
type Configuration struct {
	ServerURL      string `json:"server_url"`
	RepoPattern    string `json:"repo_pattern"`
	VersionPattern string `json:"version_pattern"`
}

/*
Config is the configuration instance
*/
var Config *Configuration

func loadConfig() {
	Config = &Configuration{
		ServerURL:      "http://localhost",
		RepoPattern:    "[^\\/?%*:|\"<>.]+",
		VersionPattern: "\\d+\\.\\d+\\.\\d+",
	}

	configData, err := ioutil.ReadFile(getCofigPath())
	if err != nil {
		storeConfig()
	}

	json.Unmarshal(configData, Config)
}

func storeConfig() error {
	configData, err := json.MarshalIndent(Config, "", " ")
	if err != nil {
		return err
	}

	return ioutil.WriteFile(getCofigPath(), configData, 0644)
}

func init() {
	loadConfig()
}
