package config

import (
	"os"
	"strconv"
)

// Config holds all configuration for the server
type Config struct {
	Port            int
	ShutdownTimeout int
	Environment     string
}

// New creates a new Config with values from environment variables
func New() *Config {
	port, _ := strconv.Atoi(getEnv("PORT", "8080"))
	shutdownTimeout, _ := strconv.Atoi(getEnv("SHUTDOWN_TIMEOUT", "10"))
	
	return &Config{
		Port:            port,
		ShutdownTimeout: shutdownTimeout,
		Environment:     getEnv("ENV", "development"),
	}
}

// getEnv gets an environment variable or returns a default value
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
} 