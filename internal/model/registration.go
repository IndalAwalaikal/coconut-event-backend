package model

import "time"

// Registration maps to `registrations` table
type Registration struct {
    ID          string    `db:"id" json:"id"`
    EventID     string    `db:"event_id" json:"eventId"`
    Name        string    `db:"name" json:"name"`
    Whatsapp    string    `db:"whatsapp" json:"whatsapp"`
    Institution string    `db:"institution" json:"institution"`
    ProofImage  string    `db:"proof_image" json:"proofImage"`
    FileName    string    `db:"file_name" json:"fileName"`
    RegisteredAt time.Time `db:"registered_at" json:"registeredAt"`
    CreatedAt   time.Time `db:"created_at" json:"createdAt"`
    UpdatedAt   time.Time `db:"updated_at" json:"updatedAt"`
}
