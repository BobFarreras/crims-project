package config

import (
	"net/http"
	"os"
	"strings"
	"time"
)

type Config struct {
	Port               string
	Environment        string
	PocketBaseURL      string
	PocketBaseTimeout  time.Duration
	AllowedOrigins     []string // NOVA CONFIGURACIÓ
	AuthCookieSecure   bool
	AuthCookieSameSite http.SameSite
}

func Load() (Config, error) {
	// Llegir orígens de l'entorn o usar per defecte localhost:3000
	rawOrigins := getEnvDefault("CORS_ALLOWED_ORIGINS", "http://localhost:3000")
	allowedOrigins := strings.Split(rawOrigins, ",")

	config := Config{
		Port:               getEnvDefault("PORT", "8080"),
		Environment:        getEnvDefault("ENVIRONMENT", "development"),
		PocketBaseURL:      os.Getenv("PB_URL"),
		PocketBaseTimeout:  5 * time.Second,
		AllowedOrigins:     allowedOrigins,
		AuthCookieSecure:   getEnvDefault("ENVIRONMENT", "development") == "production",
		AuthCookieSameSite: parseSameSite(getEnvDefault("AUTH_COOKIE_SAMESITE", "lax")),
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

func parseSameSite(value string) http.SameSite {
	switch strings.ToLower(strings.TrimSpace(value)) {
	case "none":
		return http.SameSiteNoneMode
	case "strict":
		return http.SameSiteStrictMode
	default:
		return http.SameSiteLaxMode
	}
}
