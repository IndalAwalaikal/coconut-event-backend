package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"github.com/IndalAwalaikal/coconut-event-hub/backend/internal/config"
	"github.com/IndalAwalaikal/coconut-event-hub/backend/internal/middleware"
	"github.com/IndalAwalaikal/coconut-event-hub/backend/internal/router"
)

func main() {
    // Load .env automatically in development if present
    _ = godotenv.Load()

    db, err := config.InitDB()
    if err != nil {
        log.Fatalf("failed to connect to db: %v", err)
    }
    defer db.Close()

    // Log DB connection info for debugging (do not log password in real prod)
    log.Printf("DB_HOST=%s DB_PORT=%s DB_USER=%s DB_NAME=%s", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_NAME"))

    r := router.NewRouter(db)
    // wrap with CORS middleware
    r = middleware.CORS(r)

    port := os.Getenv("PORT")
    if port == "" {
        port = "8010"
    }
    addr := ":" + port
    log.Printf("starting server on %s", addr)
    log.Fatal(http.ListenAndServe(addr, r))
}
