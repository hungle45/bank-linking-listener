package config

import (
	"io"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Server   ServerConfig   `yml:"server"`
	Database DatabaseConfig `yml:"database"`
	Kafka    KafkaConfig    `yml:"kafka"`
}

type ServerConfig struct {
	Host  string      `yml:"host"`
	Port  int         `yml:"port"`
	Admin AdminConfig `yml:"admin"`
}

type AdminConfig struct {
	Email    string `yml:"email"`
	Password string `yml:"password"`
}

type DatabaseConfig struct {
	Host     string `yml:"host"`
	Port     int    `yml:"port"`
	User     string `yml:"user"`
	Password string `yml:"password"`
	Timezone string `yml:"timezone"`
	DBName   string `yml:"dbname"`
}

type KafkaConfig struct {
	Brokers []string `yml:"brokers"`
}

func LoadConfig(configFilePath string) *Config {
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

	return &config
}
