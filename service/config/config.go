package config

import (
	"encoding/json"
	"io"
	"os"
)

type Config struct {
	Host string `json:"host"`
	Port string `json:"port"`
	Dsn  string `json:"dsn"`
}

func Parse(pathToConfig string) (*Config, error) {
	var cfg Config
	file, err := os.Open(pathToConfig)
	if err != nil {
		return nil, err
	}

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(fileBytes, &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
