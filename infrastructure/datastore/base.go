package datastore

import (
	"gopkg.in/yaml.v3"
	"io"
	"os"
	"taveler/config"
)

func GetAppConfig() (*config.AppConfig, error) {
	file, err := os.Open("config.yml")
	if err != nil {
		return nil, err
	}

	bytes, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var conf config.AppConfig
	err = yaml.Unmarshal(bytes, &conf)
	if err != nil {
		return nil, err
	}

	if err = file.Close(); err != nil {
		return nil, err
	}

	return &conf, nil
}
