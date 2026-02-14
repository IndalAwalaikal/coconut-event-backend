package response

import (
	"time"

	"github.com/IndalAwalaikal/coconut-event-hub/backend/internal/model"
)

// DocumentationResponse is the DTO returned to clients for documentation resources
type DocumentationResponse struct {
    ID            string    `json:"id"`
    EventId       string    `json:"eventId"`
    Category      string    `json:"category"`
    CategoryLabel string    `json:"categoryLabel"`
    EventTitle    string    `json:"eventTitle"`
    Year          int       `json:"year"`
    Images        []string  `json:"images"`
    Description   string    `json:"description"`
    CreatedAt     time.Time `json:"createdAt"`
    UpdatedAt     time.Time `json:"updatedAt"`
}

// FromModel maps model.Documentation to DocumentationResponse
func DocumentationFromModel(m *model.Documentation) DocumentationResponse {
    if m == nil {
        return DocumentationResponse{}
    }
    return DocumentationResponse{
        ID:            m.ID,
        EventId:       m.EventID,
        Category:      m.Category,
        CategoryLabel: m.CategoryLabel,
        EventTitle:    m.EventTitle,
        Year:          m.Year,
        Images:        m.Images,
        Description:   m.Description,
        CreatedAt:     m.CreatedAt,
        UpdatedAt:     m.UpdatedAt,
    }
}

// FromModelSlice maps a slice of model.Documentation to slice of DocumentationResponse
func DocumentationsFromModels(ms []model.Documentation) []DocumentationResponse {
    out := make([]DocumentationResponse, 0, len(ms))
    for i := range ms {
        m := ms[i]
        out = append(out, DocumentationFromModel(&m))
    }
    return out
}
