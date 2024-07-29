package config

import "os"

// Config holds the application configuration.
type Config struct {
    Port   string
    DBURI  string
    DBName string
}

// LoadConfig loads configuration from environment variables.
func LoadConfig() Config {
    return Config{
        Port:   getEnv("PORT", "8080"),
        DBURI:  getEnv("DB_URI", "mongodb://localhost:27017"),
        DBName: getEnv("DB_NAME", "GO_DB"), // Default to "testdb" if not set
    }
}

// getEnv retrieves an environment variable or returns a default value.
func getEnv(key, defaultValue string) string {
    if value, exists := os.LookupEnv(key); exists {
        return value
    }
    return defaultValue
}
