package config

import (
    "os"
)

type Config struct {
    Port  string
    DBURI string
}

func LoadConfig() Config {
    return Config{
        Port:  getEnv("PORT", "8080"),
        DBURI: getEnv("DB_URI", "mongodb://localhost:27017"),
    }
}

func getEnv(key, defaultValue string) string {
    if value, exists := os.LookupEnv(key); exists {
        return value
    }
    return defaultValue
}
