package request

// LoginRequest represents payload for admin login
type LoginRequest struct {
    Username string `json:"username"`
    Password string `json:"password"`
}
