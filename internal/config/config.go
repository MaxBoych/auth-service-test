package config

import (
	"encoding/json"
	"errors"
	"github.com/go-playground/validator/v10"
	"os"
)

const configPath = "./internal/config/config.json"

type Config struct {
	Server struct {
		Host string `validate:"required"`
	}
	DB struct {
		Type string `validate:"required"`
	}
	IsProduction bool `default:"false"`
}

func LoadConfig(config *Config) error {
	if config.IsProduction {
		return loadProdConfig(config)
	}
	return loadDevConfig(config)
}

func loadDevConfig(config *Config) error {
	jsonFile, err := os.Open(configPath)
	if err != nil {
		return err
	}

	err = json.NewDecoder(jsonFile).Decode(&config)
	if err != nil {
		return err
	}

	err = validator.New().Struct(config)
	if err != nil {
		return err
	}

	return nil
}

func loadProdConfig(_ *Config) error {
	// for example, load config from vault
	return errors.New("prod config is required")
}
