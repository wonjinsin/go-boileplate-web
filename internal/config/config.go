package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config holds all application configuration
type Config struct {
	Port string
	Env  string
}

// Load reads configuration from .env.local file and environment variables
// Environment variables take priority over file values
func Load() *Config {
	// Try to load .env.local file (ignore error if file doesn't exist)
	_ = godotenv.Load(".env.local")

	cfg := &Config{
		Port: mustGetEnv("PORT"),
		Env:  mustGetEnv("ENV"),
	}

	log.Printf("Configuration loaded: ENV=%s, PORT=%s", cfg.Env, cfg.Port)

	return cfg
}

// mustGetEnv reads an environment variable or panics if not found
func mustGetEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		panic(fmt.Sprintf("required environment variable %s is not set", key))
	}
	return value
}
