package repository

import (
	"database/sql"
	"time"

	"github.com/IndalAwalaikal/coconut-event-hub/backend/internal/model"
)

type PosterRepository interface {
    Create(p *model.Poster) error
    Update(p *model.Poster) error
    Delete(id string) error
    GetByID(id string) (*model.Poster, error)
    List() ([]model.Poster, error)
}

type posterRepository struct {
    db *sql.DB
}

func NewPosterRepository(db *sql.DB) PosterRepository {
    return &posterRepository{db: db}
}

func (r *posterRepository) Create(p *model.Poster) error {
    now := time.Now()
    if p.ID == "" {
        // caller should set UUID
    }
    _, err := r.db.Exec(`INSERT INTO posters (id, title, type, image, date, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?)`,
        p.ID, p.Title, p.Type, p.Image, p.Date, now, now)
    if err != nil {
        return err
    }
    p.CreatedAt = now
    p.UpdatedAt = now
    return nil
}

func (r *posterRepository) Update(p *model.Poster) error {
    now := time.Now()
    _, err := r.db.Exec(`UPDATE posters SET title = ?, type = ?, image = ?, date = ?, updated_at = ? WHERE id = ?`,
        p.Title, p.Type, p.Image, p.Date, now, p.ID)
    if err != nil {
        return err
    }
    p.UpdatedAt = now
    return nil
}

func (r *posterRepository) Delete(id string) error {
    _, err := r.db.Exec(`DELETE FROM posters WHERE id = ?`, id)
    return err
}

func (r *posterRepository) GetByID(id string) (*model.Poster, error) {
    row := r.db.QueryRow(`SELECT id, title, type, image, date, created_at, updated_at FROM posters WHERE id = ?`, id)
    var p model.Poster
    var createdAt, updatedAt sql.NullTime
    if err := row.Scan(&p.ID, &p.Title, &p.Type, &p.Image, &p.Date, &createdAt, &updatedAt); err != nil {
        return nil, err
    }
    if createdAt.Valid {
        p.CreatedAt = createdAt.Time
    }
    if updatedAt.Valid {
        p.UpdatedAt = updatedAt.Time
    }
    return &p, nil
}

func (r *posterRepository) List() ([]model.Poster, error) {
    rows, err := r.db.Query(`SELECT id, title, type, image, date, created_at, updated_at FROM posters ORDER BY created_at DESC`)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    out := make([]model.Poster, 0)
    for rows.Next() {
        var p model.Poster
        var createdAt, updatedAt sql.NullTime
        if err := rows.Scan(&p.ID, &p.Title, &p.Type, &p.Image, &p.Date, &createdAt, &updatedAt); err != nil {
            return nil, err
        }
        if createdAt.Valid {
            p.CreatedAt = createdAt.Time
        }
        if updatedAt.Valid {
            p.UpdatedAt = updatedAt.Time
        }
        out = append(out, p)
    }
    return out, nil
}
