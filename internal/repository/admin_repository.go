package repository

import (
	"database/sql"

	"github.com/IndalAwalaikal/coconut-event-hub/backend/internal/model"
)

type AdminRepository struct {
    DB *sql.DB
}

func NewAdminRepository(db *sql.DB) *AdminRepository {
    return &AdminRepository{DB: db}
}

func (r *AdminRepository) GetByUsername(username string) (*model.Admin, error) {
    q := `SELECT id,username,password_hash,name,role,created_at,updated_at FROM admins WHERE username = ? LIMIT 1`
    row := r.DB.QueryRow(q, username)
    var a model.Admin
    if err := row.Scan(&a.ID, &a.Username, &a.PasswordHash, &a.Name, &a.Role, &a.CreatedAt, &a.UpdatedAt); err != nil {
        return nil, err
    }
    return &a, nil
}

func (r *AdminRepository) Create(a *model.Admin) error {
    q := `INSERT INTO admins (username,password_hash,name,role,created_at,updated_at) VALUES (?,?,?,?,NOW(),NOW())`
    _, err := r.DB.Exec(q, a.Username, a.PasswordHash, a.Name, a.Role)
    return err
}
