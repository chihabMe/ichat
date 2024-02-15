package config

import (
	"os"
	"time"
)

// Config holds configuration settings for the application.
type Config struct {
    Port            string
    DBDriver        string
    DBURL           string
    AccessTokenTTL  time.Duration
    RefreshTokenTTL time.Duration
    // Add more configuration fields as needed
}

// NewConfig creates a new Config instance with default values.
func NewConfig() *Config {
    return &Config{
        Port:            getEnvOrDefault("PORT", "3000"),
        DBDriver:        getEnvOrDefault("DB_DRIVER", "sqlite3"),
        DBURL:           getEnvOrDefault("DB_URL", "sqlite.db"),
        AccessTokenTTL:  parseDuration(getEnvOrDefault("ACCESS_TOKEN_TTL", "24h")),    // Default access token TTL is 24 hours
        RefreshTokenTTL: parseDuration(getEnvOrDefault("REFRESH_TOKEN_TTL", "720h")), // Default refresh token TTL is 30 days
    }
}

// getEnvOrDefault retrieves the value of the environment variable key, or returns defaultValue if not set.
func getEnvOrDefault(key, defaultValue string) string {
    value := os.Getenv(key)
    if value == "" {
        return defaultValue
    }
    return value
}

// parseDuration parses a duration string into a time.Duration. It panics if the duration string is invalid.
func parseDuration(durationStr string) time.Duration {
    duration, err := time.ParseDuration(durationStr)
    if err != nil {
        panic(err)
    }
    return duration
}
