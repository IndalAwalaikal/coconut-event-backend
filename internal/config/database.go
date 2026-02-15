package config

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// InitDB initializes a DB connection using environment variables.
// Defaults are suitable for a local XAMPP installation.
func InitDB() (*sql.DB, error) {
    host := getenv("DB_HOST", "mysql")
    port := getenv("DB_PORT", "3306")
    user := getenv("DB_USER", "root")
    pass := getenv("DB_PASS", "password123")
    name := getenv("DB_NAME", "coconut_event_hub")

    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&charset=utf8mb4&collation=utf8mb4_general_ci",
        user, pass, host, port, name)

    db, err := sql.Open("mysql", dsn)
    if err != nil {
        return nil, err
    }

    // Basic connection tuning
    db.SetConnMaxLifetime(5 * time.Minute)
    db.SetMaxOpenConns(25)
    db.SetMaxIdleConns(25)

    if err := db.Ping(); err != nil {
        return nil, err
    }

    return db, nil
}

func getenv(key, fallback string) string {
    if v := os.Getenv(key); v != "" {
        return v
    }
    return fallback
}
