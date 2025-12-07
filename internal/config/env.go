package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config holds all application configuration.
type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
}

// ServerConfig holds server configuration.
type ServerConfig struct {
	Host string
	Port string
}

// LoadConfig loads configuration from environment variables.
func LoadConfig() Config {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Failed to load environment variables: %v", err)
	}

	return Config{
		Server: ServerConfig{
			Host: getEnv("SERVER_HOST", "localhost"),
			Port: getEnv("SERVER_PORT", "8080"),
		},
		Database: DatabaseConfig{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnv("DB_PORT", "3306"),
			User:     getEnv("DB_USER", "root"),
			Password: getEnv("DB_PASSWORD", "Kikuk@123"),
			DBName:   getEnv("DB_NAME", "bookstore"),
		},
	}
}

// getEnv gets environment variable with default value.
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
