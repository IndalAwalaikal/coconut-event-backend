package constant

// EventCategory values used across backend
type EventCategory string

const (
    CategoryOpenClass EventCategory = "open-class"
    CategoryWebinar   EventCategory = "webinar"
    CategorySeminar   EventCategory = "seminar"
    CategoryBootcamp  EventCategory = "bootcamp"
)

var CategoryLabels = map[EventCategory]string{
    CategoryOpenClass: "COCONUT Open Class",
    CategoryWebinar:   "Webinar",
    CategorySeminar:   "Seminar",
    CategoryBootcamp:  "Bootcamp",
}
