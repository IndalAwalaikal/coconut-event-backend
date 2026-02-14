package model

// EventSpeaker is a lightweight struct representing a speaker for an event.
// The main event model currently stores speaker1/2/3 fields; this struct is
// provided for DTO use and future extensibility.
type EventSpeaker struct {
    Name  string `json:"name"`
    Title string `json:"title,omitempty"`
    Bio   string `json:"bio,omitempty"`
}
