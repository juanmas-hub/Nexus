package config

import (
	"os"
)

type Config struct {
    AuthServiceURL    string
    AuthServiceTimeout string
    CatalogServiceURL string
    Port              string
}

func Load() *Config {
    return &Config{
        AuthServiceURL:    getEnv("AUTH_SERVICE_URL", "http://localhost:8081"),
        AuthServiceTimeout: getEnv("AUTH_SERVICE_TIMEOUT", "60s"),
        CatalogServiceURL: getEnv("CATALOG_SERVICE_URL", "http://localhost:8082"),
        Port:              getEnv("PORT", "8080"),
    }
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}