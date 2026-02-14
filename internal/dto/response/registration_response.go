package response

import "time"

type RegistrationResponse struct {
    ID          string    `json:"id"`
    EventID     string    `json:"eventId"`
    Name        string    `json:"name"`
    Whatsapp    string    `json:"whatsapp"`
    Institution string    `json:"institution"`
    ProofImage  string    `json:"proofImage"`
    FileName    string    `json:"fileName"`
    RegisteredAt time.Time `json:"registeredAt"`
}
