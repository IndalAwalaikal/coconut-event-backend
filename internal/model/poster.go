package model

import "time"

// Poster represents a single poster for previous activities
type Poster struct {
    ID        string    `db:"id" json:"id"`
    Title     string    `db:"title" json:"title"`
    Type      string    `db:"type" json:"type"`
    Image     string    `db:"image" json:"image"`
    Date      string    `db:"date" json:"date"`
    CreatedAt time.Time `db:"created_at" json:"createdAt"`
    UpdatedAt time.Time `db:"updated_at" json:"updatedAt"`
}
