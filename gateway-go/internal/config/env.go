package config

import (
    "log"
    "os"

    "github.com/joho/godotenv"
)

// LoadEnv reads the .env file.
// It doesn't crash if file is missing (useful for production), 
// but warns us in development.
func LoadEnv() {
    err := godotenv.Load()
    if err != nil {
        log.Println("Warning: No .env file found, using system defaults")
    }
}

func Get(key string, fallback string) string {
    if value, exists := os.LookupEnv(key); exists {
        return value
    }
    return fallback
}
