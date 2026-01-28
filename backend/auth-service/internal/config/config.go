package config

import(
	"os"
	"github.com/joho/godotenv"
)

type Config struct {
    APP_MODE    	string
    DATABASE_URL 	string
	JWT_SECRET 		string
    Port            string
}

func Load() *Config {
	_ = godotenv.Load()

	return &Config{
		APP_MODE:		os.Getenv("APP_MODE"),
		DATABASE_URL:	os.Getenv("DATABASE_URL"),
		JWT_SECRET:		os.Getenv("JWT_SECRET"),
		Port:			getEnv("PORT", "8081"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}