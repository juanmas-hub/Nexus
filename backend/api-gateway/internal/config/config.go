package config

import (
	"os"
    "time"
	"github.com/joho/godotenv"
)

type Config struct {
    AuthServiceURL    	string
    AuthServiceTimeout 	time.Duration
    CatalogServiceURL 	string

	IsProd				bool		
    AllowedOrigins 		[]string
    Port              	string
}

func Load() *Config {
	_ = godotenv.Load()

	return &Config{
		Port:               getEnv("PORT", "8080"),
		AuthServiceURL:     os.Getenv("AUTH_SERVICE_URL"),
		AuthServiceTimeout: parseDurationEnv("AUTH_SERVICE_TIMEOUT", "30s"),
		CatalogServiceURL:  os.Getenv("CATALOG_SERVICE_URL"),
		AllowedOrigins:     []string{"http://localhost:5173", "https://nexus-b6b.pages.dev"},
		IsProd: 			os.Getenv("APP_MODE") == "prod",
	}
}

func parseDurationEnv(key string, defaultDuration string) time.Duration {
	value := os.Getenv(key)
	if value == "" {
		value = defaultDuration
	}

	duration, err := time.ParseDuration(value)
	if err != nil {
		fallback, _ := time.ParseDuration(defaultDuration)
		return fallback
	}
	return duration
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}