package core

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// Config reads an environment variable value
func Config(key string) string {
	value := os.Getenv(key)
	if value != "" {
		return value
	}

	// Fallback to loading from .env file in development
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf("Error loading .env file: %s\n", err)
		return "" // or handle the error according to your needs
	}

	value = os.Getenv(key)
	if value == "" {
		fmt.Printf("Warning: Environment variable %s is not set\n", key)
	}
	return value
}
