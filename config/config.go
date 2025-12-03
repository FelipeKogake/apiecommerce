package config

import (
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	Port string
	Database DatabaseConfig
}

type DatabaseConfig struct {
	URI string
	Nome string
}

func Load() (*Config, error) {
	godotenv.Load("../../.env")
	config := &Config{
		Port: os.Getenv("PORT"),
		Database: DatabaseConfig{
			URI: os.Getenv("URI"),
			Nome: os.Getenv("DATABASE"),
		},
	}
	return config, nil
}