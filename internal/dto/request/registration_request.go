package request

// CreateRegistrationRequest represents the payload expected when a user registers
type CreateRegistrationRequest struct {
    EventID     string `json:"eventId" form:"event_id"`
    Name        string `json:"name" form:"name"`
    Whatsapp    string `json:"whatsapp" form:"whatsapp"`
    Institution string `json:"institution" form:"institution"`
    // If the client uploads a file, the server will handle it as multipart form; file name is optional
    FileName    string `json:"fileName" form:"file_name"`
}
