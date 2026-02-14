package public

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/IndalAwalaikal/coconut-event-hub/backend/internal/service"
)

type EventController struct {
    Service *service.EventService
}

func NewEventController(s *service.EventService) *EventController {
    return &EventController{Service: s}
}

// GET /api/events
func (c *EventController) List(w http.ResponseWriter, r *http.Request) {
    // Ambil dari "search" atau fallback ke "q"
    search := r.URL.Query().Get("search")
    if search == "" {
        search = r.URL.Query().Get("q") // fallback
    }
    category := r.URL.Query().Get("category")
    
    log.Printf("Fetching events: category=%s, search=%s", category, search)
    
    events, err := c.Service.List(category, search)
    if err != nil {
        log.Printf("list events err: %v", err)
        http.Error(w, "internal", http.StatusInternalServerError)
        return
    }
    
    log.Printf("Found %d events", len(events))
    
    w.Header().Set("Content-Type", "application/json")
    _ = json.NewEncoder(w).Encode(events)
}

// GET /api/events/{id}
func (c *EventController) Get(w http.ResponseWriter, r *http.Request) {
    // mux Vars not imported here in public controller to keep it simple; use query param id
    id := r.URL.Query().Get("id")
    if id == "" {
        http.Error(w, "id required", http.StatusBadRequest)
        return
    }
    ev, err := c.Service.GetByID(id)
    if err != nil {
        log.Printf("get event err: %v", err)
        http.Error(w, "not found", http.StatusNotFound)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    _ = json.NewEncoder(w).Encode(ev)
}
