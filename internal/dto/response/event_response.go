package response

import "time"

type EventListItem struct {
	ID            string    `json:"id"`
	Category      string    `json:"category"`
	CategoryLabel string    `json:"categoryLabel"`
	Title         string    `json:"title"`
	Description   string    `json:"description"`
	Date          *time.Time `json:"date"`
	Time          string    `json:"time"`
	Location      string    `json:"location"`
	Quota         int       `json:"quota"`
	Registered    int       `json:"registered"`
	Poster        string    `json:"poster"`
	EventType     string    `json:"eventType"` 
	Price         int       `json:"price"`     
	Available     bool      `json:"available"`
}