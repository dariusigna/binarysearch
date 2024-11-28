package config

import (
	"fmt"
	log "log/slog"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Port              string `env:"PORT" default:"8080"`
	LogLevel          string `env:"LOG_LEVEL" default:"info"`
	InputFile         string `env:"INPUT_FILE" default:"input.txt"`
	ConformationLevel int    `env:"CONFORMATION_LEVEL" default:"10"`
}

func LoadConfig() (*Config, error) {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Warn("No .env file found, using default values.")
	}

	var cfg Config
	if err = envconfig.Process("", &cfg); err != nil {
		return nil, fmt.Errorf("failed to load configuration: %v", err)
	}
	return &cfg, nil
}
