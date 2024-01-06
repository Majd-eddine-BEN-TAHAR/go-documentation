package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config struct to hold application configuration
type Config struct {
    DatabasePath string
    JWTSecret    string
}

// LoadConfig loads application configurations from .env file
func LoadConfig() *Config {
    // Load environment variables from .env file
    err := godotenv.Load()
    fmt.Println(err)
    if err != nil {
        log.Fatal("Error loading .env file")
    }
    // Return the configurations as a Config struct
    return &Config{
        DatabasePath: os.Getenv("DATABASE_PATH"), // Database path from .env
        JWTSecret:    os.Getenv("JWT_SECRET"),    // JWT secret key from .env
    }
}