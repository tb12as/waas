package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppEnv  string
	AppPort string
	// OriginHost is the registrable domain (e.g. "syfq.my.id"). Any origin on
	// this host or a subdomain of it is allowed for CORS. Empty disables CORS.
	OriginHost string

	DbName     string
	DbUsername string
	DbPassword string
	DbPort     string
	DbHost     string
}

func Load() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	cfg := &Config{
		AppEnv:  getEnv("APP_ENV", "local"),
		AppPort: getEnv("APP_PORT", "8080"),

		OriginHost: getEnv("ORIGIN_HOST", "syfq.my.id"),

		DbName:     getEnv("DB_NAME", "waas"),
		DbUsername: getEnv("DB_USERNAME", "root"),
		DbPassword: getEnv("DB_PASSWORD", "root"),
		DbHost:     getEnv("DB_HOST", "localhost"),
		DbPort:     getEnv("DB_PORT", "3306"),
	}

	return cfg, nil
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
