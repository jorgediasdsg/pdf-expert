package config

import (
	"os"
)

type Config struct {
	Env        string
	HTTPPort   string
	TempFolder string
}

func Load() Config {
	cfg := Config{
		Env:        get("APP_ENV", "dev"),
		HTTPPort:   get("HTTP_PORT", "8080"),
		TempFolder: get("TEMP_FOLDER", "./tmp"),
	}

	return cfg
}

func get(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
