package config

import (
	"os"
)

type Config struct {
	AuthServiceURL string
	Port           string
}

func Load() *Config {
	return &Config{
		AuthServiceURL: getEnv("AUTH_SERVICE_URL", "http://localhost:8081"),
		Port:           getEnv("PORT", "8080"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}