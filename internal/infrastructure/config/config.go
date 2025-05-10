package config

import (
	"log"
	"os"
)

type Config struct {
	Env         string
	Port        string
	DatabaseURL string
	SupaBaseURL string
	SupaBaseKey string
}

func Load() *Config {
	cfg := &Config{
		Env:         getEnv("ENV", "development"),
		Port:        getEnv("PORT", "8080"),
		DatabaseURL: getEnv("DATABASE_URL", ""),
		SupaBaseURL: getEnv("SUPABASE_URL", ""),
		SupaBaseKey: getEnv("SUPABASE_KEY", ""),
	}

	if cfg.DatabaseURL == "" {
		log.Fatal("Database URL is required")
	}

	return cfg
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
