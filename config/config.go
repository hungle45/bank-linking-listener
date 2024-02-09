package config

import (
	"io"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Server serverConfig `yml:"app"`
}

type serverConfig struct {
	Host string `yml:"host"`
	Port int    `yml:"port"`
}

func LoadConfig(configFilePath string) Config {
	file, err := os.Open(configFilePath)
	if err != nil {
		log.Fatalf("Error opening config file: %s", err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("Error reading config file: %s", err)
	}

	var config Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		log.Fatalf("Error unmarshaling config: %s", err)
	}

	return config
}
