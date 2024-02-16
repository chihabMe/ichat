package config

import (
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
    Port            string
    DBDriver        string
    DBUser          string
    DBPassword      string
    DBHost          string
    DBName          string
    AccessTokenTTL  time.Duration
    RefreshTokenTTL time.Duration
}

func InitConfig() *Config {
    return &Config{
        Port:            GetEnvOrDefault("PORT", "3000"),
        DBDriver:        GetEnvOrDefault("DB_DRIVER", "mysql"),
        DBUser:          GetEnvOrDefault("DB_USER", "root"),
        DBPassword:      GetEnvOrDefault("DB_PASSWORD", ""),
        DBHost:          GetEnvOrDefault("DB_HOST", "localhost"),
        DBName:          GetEnvOrDefault("DB_NAME", "dbname"),
        AccessTokenTTL:  parseDuration(GetEnvOrDefault("ACCESS_TOKEN_TTL", "24h")),
        RefreshTokenTTL: parseDuration(GetEnvOrDefault("REFRESH_TOKEN_TTL", "720h")),
        // Add default values for other configuration fields
    }
}

func (c *Config) DatabaseDSN() string {
    return fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", c.DBUser, c.DBPassword, c.DBHost, c.DBName)
}

// getEnvOrDefault retrieves the value of the environment variable key, or returns defaultValue if not set.
func GetEnvOrDefault(key, defaultValue string) string {
    value := os.Getenv(key)
    if value == "" {
        return defaultValue
    }
    return value
}
func InitDotenv()error{
    return godotenv.Load()
}

func parseDuration(durationStr string) time.Duration {
    duration, err := time.ParseDuration(durationStr)
    if err != nil {
        panic(fmt.Sprintf("Error parsing duration: %s", err))
    }
    return duration
}
