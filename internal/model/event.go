package model

import "time"

// Event maps to `events` table
type Event struct {
    ID            string    `db:"id" json:"id"`
    Category      string    `db:"category" json:"category"`
    CategoryLabel string    `db:"category_label" json:"categoryLabel"`
    Title         string    `db:"title" json:"title"`
    Description   string    `db:"description" json:"description"`
    Rules         []string  `db:"rules" json:"rules"`
    Benefits      []string  `db:"benefits" json:"benefits"`
    Date          *time.Time `db:"date" json:"date"`
    Time          string    `db:"time" json:"time"`
    Speaker1      string    `db:"speaker1" json:"speaker1"`
    Speaker2      string    `db:"speaker2" json:"speaker2"`
    Speaker3      string    `db:"speaker3" json:"speaker3"`
    Moderator     string    `db:"moderator" json:"moderator"`
    Location      string    `db:"location" json:"location"`
    Quota         int       `db:"quota" json:"quota"`
    Registered    int       `db:"registered" json:"registered"`
    Poster        string    `db:"poster" json:"poster"`
    EventType     string    `db:"event_type" json:"eventType"` 
    Price         int       `db:"price" json:"price"`         
    Available     bool      `db:"available" json:"available"`
    CreatedAt     time.Time `db:"created_at" json:"createdAt"`
    UpdatedAt     time.Time `db:"updated_at" json:"updatedAt"`
}