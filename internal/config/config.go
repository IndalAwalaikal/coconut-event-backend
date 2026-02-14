package config

import (
	"os"
)

type AppConfig struct {
    DBHost string
    DBPort string
    DBUser string
    DBPass string
    DBName string
    Port   string
}

// LoadConfig reads configuration from environment variables (defaults provided)
func LoadConfig() *AppConfig {
    dbHost := os.Getenv("DB_HOST")
    if dbHost == "" { dbHost = "127.0.0.1" }
    dbPort := os.Getenv("DB_PORT")
    if dbPort == "" { dbPort = "3306" }
    dbUser := os.Getenv("DB_USER")
    if dbUser == "" { dbUser = "root" }
    dbPass := os.Getenv("DB_PASS")
    if dbPass == "" { dbPass = "password123" }
    dbName := os.Getenv("DB_NAME")
    if dbName == "" { dbName = "coconut_event_hub" }
    port := os.Getenv("PORT")
    if port == "" { port = "8080" }
    return &AppConfig{
        DBHost: dbHost,
        DBPort: dbPort,
        DBUser: dbUser,
        DBPass: dbPass,
        DBName: dbName,
        Port:   port,
    }
}

