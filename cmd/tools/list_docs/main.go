package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"

	"github.com/IndalAwalaikal/coconut-event-hub/backend/internal/config"
	"github.com/IndalAwalaikal/coconut-event-hub/backend/internal/repository"
)

func main() {
    _ = godotenv.Load()
    db, err := config.InitDB()
    if err != nil {
        log.Fatalf("db init: %v", err)
    }
    defer db.Close()
    repo := repository.NewDocumentationRepository(db)
    docs, err := repo.List("", 0, "", 0, 0)
    if err != nil {
        log.Fatalf("list err: %v", err)
    }
    fmt.Printf("found %d docs\n", len(docs))
    for _, d := range docs {
        fmt.Printf("%s %s %d %v\n", d.ID, d.EventTitle, d.Year, d.Images)
    }
}

