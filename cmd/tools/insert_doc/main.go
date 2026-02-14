package main

import (
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/joho/godotenv"

	"github.com/IndalAwalaikal/coconut-event-hub/backend/internal/config"
	"github.com/IndalAwalaikal/coconut-event-hub/backend/internal/model"
	"github.com/IndalAwalaikal/coconut-event-hub/backend/internal/repository"
)

func main() {
    // load .env so this tool uses same DB settings as server
    _ = godotenv.Load()

    db, err := config.InitDB()
    if err != nil {
        log.Fatalf("failed to connect db: %v", err)
    }
    defer db.Close()

    repo := repository.NewDocumentationRepository(db)
    doc := &model.Documentation{
        ID:            uuid.New().String(),
        EventID:       "",
        Category:      "webinar",
        CategoryLabel: "Webinar",
        EventTitle:    "Test Insert Doc",
        Year:          2026,
        Images:        []string{"/storage/documentations/test1.jpg"},
        Description:   "Inserted by tool",
    }
    if err := repo.Create(doc); err != nil {
        log.Fatalf("create failed: %v", err)
    }
    fmt.Printf("inserted doc id=%s\n", doc.ID)
}

