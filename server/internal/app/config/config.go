package config

import (
	"fmt"
	"os"
)

// Config holds configuration settings for the application.
type Config struct {
    Port            string
    DBDriver        string
    DBUser          string
    DBPassword      string
    DBHost          string
    DBName          string
    AccessTokenTTL  string // Change this to the appropriate type for your token TTL
    RefreshTokenTTL string // Change this to the appropriate type for your token TTL
    // Add more configuration fields as needed
}

// NewConfig creates a new Config instance with default values.
func InitConfig() *Config {
    return &Config{
        Port:            getEnvOrDefault("PORT", "3000"),
        DBDriver:        getEnvOrDefault("DB_DRIVER", "mysql"),
        DBUser:          getEnvOrDefault("DB_USER", "root"),
        DBPassword:      getEnvOrDefault("DB_PASSWORD", ""),
        DBHost:          getEnvOrDefault("DB_HOST", "localhost"),
        DBName:          getEnvOrDefault("DB_NAME", "dbname"),
        AccessTokenTTL:  getEnvOrDefault("ACCESS_TOKEN_TTL", "24h"),
        RefreshTokenTTL: getEnvOrDefault("REFRESH_TOKEN_TTL", "720h"),
        // Add default values for other configuration fields
    }
}

// DatabaseDSN returns the data source name (DSN) for the database connection.
func (c *Config) DatabaseDSN() string {
    return fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", c.DBUser, c.DBPassword, c.DBHost, c.DBName)
}

// getEnvOrDefault retrieves the value of the environment variable key, or returns defaultValue if not set.
func getEnvOrDefault(key, defaultValue string) string {
    value := os.Getenv(key)
    if value == "" {
        return defaultValue
    }
    return value
}
