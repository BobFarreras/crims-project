package config

import (
	"os"
	"time"
)

type Config struct {
	Port              string
	Environment       string
	PocketBaseURL     string
	PocketBaseTimeout time.Duration
}

func Load() (Config, error) {
	config := Config{
		Port:              getEnvDefault("PORT", "8080"),
		Environment:       getEnvDefault("ENVIRONMENT", "development"),
		PocketBaseURL:     os.Getenv("PB_URL"),
		PocketBaseTimeout: 5 * time.Second,
	}

	if rawTimeout := os.Getenv("PB_TIMEOUT"); rawTimeout != "" {
		parsed, err := time.ParseDuration(rawTimeout)
		if err != nil {
			return Config{}, err
		}
		config.PocketBaseTimeout = parsed
	}

	return config, nil
}

func getEnvDefault(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}
