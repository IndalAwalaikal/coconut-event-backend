package response

import (
	"time"

	"github.com/IndalAwalaikal/coconut-event-hub/backend/internal/model"
)

type PosterResponse struct {
    ID        string    `json:"id"`
    Title     string    `json:"title"`
    Type      string    `json:"type"`
    Image     string    `json:"image"`
    Date      string    `json:"date"`
    CreatedAt time.Time `json:"createdAt"`
    UpdatedAt time.Time `json:"updatedAt"`
}

func PosterFromModel(m *model.Poster) PosterResponse {
    if m == nil {
        return PosterResponse{}
    }
    return PosterResponse{
        ID: m.ID,
        Title: m.Title,
        Type: m.Type,
        Image: m.Image,
        Date: m.Date,
        CreatedAt: m.CreatedAt,
        UpdatedAt: m.UpdatedAt,
    }
}

func PostersFromModels(ms []model.Poster) []PosterResponse {
    out := make([]PosterResponse, 0, len(ms))
    for i := range ms {
        m := ms[i]
        out = append(out, PosterFromModel(&m))
    }
    return out
}
