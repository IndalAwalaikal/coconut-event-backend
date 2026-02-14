package response

// PaginationResponse is a generic wrapper for paginated endpoints
type PaginationResponse struct {
	Items      interface{} `json:"items"`
	Page       int         `json:"page"`
	Limit      int         `json:"limit"`
	Total      int         `json:"total"`
	TotalPages int         `json:"totalPages"`
}

