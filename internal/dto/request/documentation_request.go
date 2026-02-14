package request

// Note: admin create/update currently accept multipart/form-data. This DTO
// is provided for future JSON-based endpoints or documentation purposes.
type DocumentationCreateRequest struct {
    Category      string   `json:"category"`
    CategoryLabel string   `json:"categoryLabel"`
    EventTitle    string   `json:"eventTitle"`
    Year          int      `json:"year"`
    Images        []string `json:"images"`
    Description   string   `json:"description"`
    EventId       string   `json:"eventId"`
}

type DocumentationUpdateRequest struct {
    Category      *string  `json:"category,omitempty"`
    CategoryLabel *string  `json:"categoryLabel,omitempty"`
    EventTitle    *string  `json:"eventTitle,omitempty"`
    Year          *int     `json:"year,omitempty"`
    Images        []string `json:"images,omitempty"`
    Description   *string  `json:"description,omitempty"`
    EventId       *string  `json:"eventId,omitempty"`
}
