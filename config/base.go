package config

import (
	"gopkg.in/yaml.v3"
	"io"
	"os"
)

func GetAppConfig() (*AppConfig, error) {
	file, err := os.Open("config.yml")
	if err != nil {
		return nil, err
	}

	bytes, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var conf AppConfig
	err = yaml.Unmarshal(bytes, &conf)
	if err != nil {
		return nil, err
	}

	if err = file.Close(); err != nil {
		return nil, err
	}

	return &conf, nil
}
