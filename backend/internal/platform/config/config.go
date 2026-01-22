package config

import (
	"os"
	"strings"
	"time"
)

type Config struct {
	Port              string
	Environment       string
	PocketBaseURL     string
	PocketBaseTimeout time.Duration
	AllowedOrigins    []string // NOVA CONFIGURACIÓ
}

func Load() (Config, error) {
	// Llegir orígens de l'entorn o usar per defecte localhost:3000
	rawOrigins := getEnvDefault("CORS_ALLOWED_ORIGINS", "http://localhost:3000")
	allowedOrigins := strings.Split(rawOrigins, ",")

	config := Config{
		Port:              getEnvDefault("PORT", "8080"),
		Environment:       getEnvDefault("ENVIRONMENT", "development"),
		PocketBaseURL:     os.Getenv("PB_URL"),
		PocketBaseTimeout: 5 * time.Second,
		AllowedOrigins:    allowedOrigins,
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
