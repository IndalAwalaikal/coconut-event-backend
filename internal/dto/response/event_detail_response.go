package response

import "time"

type EventDetail struct {
	ID            string    `json:"id"`
	Category      string    `json:"category"`
	CategoryLabel string    `json:"categoryLabel"`
	Title         string    `json:"title"`
	Description   string    `json:"description"`
	Rules         []string  `json:"rules"`
	Benefits      []string  `json:"benefits"`
	Date          *time.Time `json:"date"`
	Time          string    `json:"time"`
	Speaker1      string    `json:"speaker1"`
	Speaker2      string    `json:"speaker2"`
	Speaker3      string    `json:"speaker3"`
	Moderator     string    `json:"moderator"`
	Location      string    `json:"location"`
	Quota         int       `json:"quota"`
	Registered    int       `json:"registered"`
	Poster        string    `json:"poster"`
	EventType     string    `json:"eventType"` 
	Price         int       `json:"price"`     
	Available     bool      `json:"available"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
}