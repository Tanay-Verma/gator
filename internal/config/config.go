package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DbURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func Read() (Config, error) {
	var gatorconfig Config

	fullPath, err := getConfigFilePath()
	if err != nil {
		return gatorconfig, err
	}
	data, err := os.ReadFile(fullPath)
	if err != nil {
		return gatorconfig, err
	}

	if err := json.Unmarshal(data, &gatorconfig); err != nil {
		return gatorconfig, nil
	}

	return gatorconfig, nil
}

func (c *Config) SetUser(username string) error {
	c.CurrentUserName = username

	data, err := json.Marshal(c)
	if err != nil {
		return err
	}

	fullPath, err := getConfigFilePath()
	if err != nil {
		return err
	}

	err = os.WriteFile(fullPath, data, os.ModeAppend)
	if err != nil {
		return err
	}

	return nil
}

func getConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	fullPath := filepath.Join(homeDir, configFileName)
	return fullPath, nil
}
