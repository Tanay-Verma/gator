package config

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
)

const confiFileName = ".gatorconfig.json"

type Config struct {
	DbURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func Read() Config {
	fullPath := getConfigFilePath()
	data, err := os.ReadFile(fullPath)
	if err != nil {
		log.Fatal("Error in reading gatorconfig.json Error:", err)
	}

	var gatorconfig Config
	if err := json.Unmarshal(data, &gatorconfig); err != nil {
		log.Fatal("Error in parsing gatorconfig.json Error:", err)
	}

	return gatorconfig
}

func (c *Config) SetUser(username string) {
	c.CurrentUserName = username
	data, err := json.Marshal(c)
	if err != nil {
		log.Fatal("Error in marshaling the config. Error:", err)
	}

	fullPath := getConfigFilePath()

	err = os.WriteFile(fullPath, data, os.ModeAppend)
	if err != nil {
		log.Fatal("Error in writing to gatorconfig.json file. Error:", err)
	}
}

func getConfigFilePath() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal("Error in reading user home directory. Error:", err)
	}

	fullPath := filepath.Join(homeDir, confiFileName)
	return fullPath
}
