package repository

import (
	"database/sql"

	"github.com/IndalAwalaikal/coconut-event-hub/backend/internal/model"
)

type RegistrationRepository struct {
    DB *sql.DB
}

func NewRegistrationRepository(db *sql.DB) *RegistrationRepository {
    return &RegistrationRepository{DB: db}
}

func (r *RegistrationRepository) Create(reg *model.Registration) error {
    query := `INSERT INTO registrations (id,event_id,name,whatsapp,institution,proof_image,file_name,registered_at,created_at,updated_at) VALUES (?,?,?,?,?,?,?,?,NOW(),NOW())`
    _, err := r.DB.Exec(query, reg.ID, reg.EventID, reg.Name, reg.Whatsapp, reg.Institution, reg.ProofImage, reg.FileName, reg.RegisteredAt)
    return err
}

func (r *RegistrationRepository) GetByID(id string) (*model.Registration, error) {
    q := `SELECT id,event_id,name,whatsapp,institution,proof_image,file_name,registered_at,created_at,updated_at FROM registrations WHERE id = ? LIMIT 1`
    row := r.DB.QueryRow(q, id)
    var reg model.Registration
    var registeredAt, createdAt, updatedAt sql.NullTime
    if err := row.Scan(&reg.ID, &reg.EventID, &reg.Name, &reg.Whatsapp, &reg.Institution, &reg.ProofImage, &reg.FileName, &registeredAt, &createdAt, &updatedAt); err != nil {
        return nil, err
    }
    if registeredAt.Valid {
        reg.RegisteredAt = registeredAt.Time
    }
    if createdAt.Valid {
        reg.CreatedAt = createdAt.Time
    }
    if updatedAt.Valid {
        reg.UpdatedAt = updatedAt.Time
    }
    return &reg, nil
}

func (r *RegistrationRepository) ListByEvent(eventID string) ([]model.Registration, error) {
    q := `SELECT id,event_id,name,whatsapp,institution,proof_image,file_name,registered_at,created_at,updated_at FROM registrations WHERE event_id = ? ORDER BY registered_at DESC` 
    rows, err := r.DB.Query(q, eventID)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    out := []model.Registration{}
    for rows.Next() {
        var reg model.Registration
        var registeredAt, createdAt, updatedAt sql.NullTime
        if err := rows.Scan(&reg.ID, &reg.EventID, &reg.Name, &reg.Whatsapp, &reg.Institution, &reg.ProofImage, &reg.FileName, &registeredAt, &createdAt, &updatedAt); err != nil {
            return nil, err
        }
        if registeredAt.Valid { reg.RegisteredAt = registeredAt.Time }
        if createdAt.Valid { reg.CreatedAt = createdAt.Time }
        if updatedAt.Valid { reg.UpdatedAt = updatedAt.Time }
        out = append(out, reg)
    }
    return out, nil
}
