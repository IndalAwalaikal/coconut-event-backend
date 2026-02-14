package request

// CreateEventRequest represents payload to create an event
type CreateEventRequest struct {
	ID            string   `json:"id" form:"id"`
	Category      string   `json:"category" form:"category"`
	CategoryLabel string   `json:"categoryLabel" form:"category_label"`
	Title         string   `json:"title" form:"title"`
	Description   string   `json:"description" form:"description"`
	Rules         []string `json:"rules" form:"rules"`
	Benefits      []string `json:"benefits" form:"benefits"`
	Date          string   `json:"date" form:"date"`
	Time          string   `json:"time" form:"time"`
	Speaker1      string   `json:"speaker1" form:"speaker1"`
	Speaker2      string   `json:"speaker2" form:"speaker2"`
	Speaker3      string   `json:"speaker3" form:"speaker3"`
	Moderator     string   `json:"moderator" form:"moderator"`
	Location      string   `json:"location" form:"location"`
	Quota         int      `json:"quota" form:"quota"`
	EventType     string   `json:"eventType" form:"event_type"` 
	Price         int      `json:"price" form:"price"`         
	Poster        string   `json:"poster" form:"poster"`
	Available     bool     `json:"available" form:"available"`
}