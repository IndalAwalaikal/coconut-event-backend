package model

import "time"

// Documentation represents a past event/gallery
type Documentation struct {
	ID            string    `db:"id" json:"id"`
	EventID       string    `db:"event_id" json:"eventId"`
	Category      string    `db:"category" json:"category"`
	CategoryLabel string    `db:"category_label" json:"categoryLabel"`
	EventTitle    string    `db:"event_title" json:"eventTitle"`
	Year          int       `db:"year" json:"year"`
	Images        []string  `db:"images" json:"images"`
	Description   string    `db:"description" json:"description"`
	CreatedAt     time.Time `db:"created_at" json:"createdAt"`
	UpdatedAt     time.Time `db:"updated_at" json:"updatedAt"`
}
