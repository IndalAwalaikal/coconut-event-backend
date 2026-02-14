package repository

import (
	"database/sql"
	"encoding/json"
	"strings"

	"github.com/IndalAwalaikal/coconut-event-hub/backend/internal/model"
)

// DocumentationRepository provides CRUD for documentations
type DocumentationRepository struct {
	DB *sql.DB
}

func NewDocumentationRepository(db *sql.DB) *DocumentationRepository {
	return &DocumentationRepository{DB: db}
}

// normalizeImagePath attempts to convert various stored image path formats
// (absolute filesystem paths, relative paths, or already-correct '/storage/...' paths)
// into a canonical '/storage/...' path that the backend file server exposes.
func normalizeImagePath(s string) string {
	s = strings.TrimSpace(s)
	if s == "" {
		return s
	}
	// already correct
	if strings.HasPrefix(s, "/storage/") {
		return s
	}
	// relative without leading slash
	if strings.HasPrefix(s, "storage/") {
		return "/" + s
	}
	// if an absolute filesystem path contains '/storage/' segment, cut to that segment
	if idx := strings.Index(s, "/storage/"); idx != -1 {
		return s[idx:]
	}
	// if contains 'storage/' without leading slash
	if idx := strings.Index(s, "storage/"); idx != -1 {
		return "/" + s[idx:]
	}
	// otherwise leave as-is (best-effort)
	return s
}

// internal/repository/documentation.go
func (r *DocumentationRepository) Create(d *model.Documentation) error {
    q := `INSERT INTO documentations (id,event_id,category,category_label,event_title,year,images,description,created_at,updated_at) VALUES (?,?,?,?,?,?,?,?,NOW(),NOW())`
    imgs, _ := json.Marshal(d.Images)
    
    // ✅ Handle empty event_id (convert to NULL)
    var eventID interface{}
    if d.EventID == "" {
        eventID = nil
    } else {
        eventID = d.EventID
    }
    
    _, err := r.DB.Exec(q, d.ID, eventID, d.Category, d.CategoryLabel, d.EventTitle, d.Year, string(imgs), d.Description)
    return err
}

func (r *DocumentationRepository) Update(d *model.Documentation) error {
    q := `UPDATE documentations SET event_id=?, category=?, category_label=?, event_title=?, year=?, images=?, description=?, updated_at=NOW() WHERE id=?`
    imgs, _ := json.Marshal(d.Images)
    
    // ✅ Handle empty event_id (convert to NULL)
    var eventID interface{}
    if d.EventID == "" {
        eventID = nil
    } else {
        eventID = d.EventID
    }
    
    _, err := r.DB.Exec(q, eventID, d.Category, d.CategoryLabel, d.EventTitle, d.Year, string(imgs), d.Description, d.ID)
    return err
}

func (r *DocumentationRepository) Delete(id string) error {
	q := `DELETE FROM documentations WHERE id = ?`
	_, err := r.DB.Exec(q, id)
	return err
}

func (r *DocumentationRepository) GetByID(id string) (*model.Documentation, error) {
	q := `SELECT id,event_id,category,category_label,event_title,year,images,description,created_at,updated_at FROM documentations WHERE id = ? LIMIT 1`
	row := r.DB.QueryRow(q, id)
	var d model.Documentation
	var imagesRaw sql.NullString
	var eventIDRaw sql.NullString
	var yearRaw sql.NullInt64
	if err := row.Scan(&d.ID, &eventIDRaw, &d.Category, &d.CategoryLabel, &d.EventTitle, &yearRaw, &imagesRaw, &d.Description, &d.CreatedAt, &d.UpdatedAt); err != nil {
		return nil, err
	}
	if eventIDRaw.Valid {
		d.EventID = eventIDRaw.String
	} else {
		d.EventID = ""
	}
	if yearRaw.Valid {
		d.Year = int(yearRaw.Int64)
	} else {
		d.Year = 0
	}
	if imagesRaw.Valid && strings.TrimSpace(imagesRaw.String) != "" {
		var imgs []string
		if err := json.Unmarshal([]byte(imagesRaw.String), &imgs); err == nil {
			// Normalize image paths so frontend can reliably request them from /storage/...
			for i, im := range imgs {
				imgs[i] = normalizeImagePath(im)
			}
			d.Images = imgs
		} else {
			d.Images = []string{}
		}
	} else {
		// ensure a non-nil slice so callers don't have to guard against null
		d.Images = []string{}
	}
	return &d, nil
}

func (r *DocumentationRepository) List(category string, year int, query string, limit, offset int) ([]model.Documentation, error) {
	var parts []string
	args := []interface{}{}
	base := `SELECT id,event_id,category,category_label,event_title,year,images,description,created_at,updated_at FROM documentations`
	if category != "" {
		parts = append(parts, "category = ?")
		args = append(args, category)
	}
	if year != 0 {
		parts = append(parts, "year = ?")
		args = append(args, year)
	}
	if query != "" {
		parts = append(parts, "(event_title LIKE ? OR description LIKE ?)")
		qv := "%" + query + "%"
		args = append(args, qv, qv)
	}
	if len(parts) > 0 {
		base = base + " WHERE " + strings.Join(parts, " AND ")
	}
	base = base + " ORDER BY year DESC, created_at DESC"
	if limit > 0 {
		base = base + " LIMIT ? OFFSET ?"
		args = append(args, limit, offset)
	}
	rows, err := r.DB.Query(base, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	res := []model.Documentation{}
	for rows.Next() {
		var d model.Documentation
		var imagesRaw sql.NullString
		var eventIDRaw sql.NullString
		var yearRaw sql.NullInt64
		if err := rows.Scan(&d.ID, &eventIDRaw, &d.Category, &d.CategoryLabel, &d.EventTitle, &yearRaw, &imagesRaw, &d.Description, &d.CreatedAt, &d.UpdatedAt); err != nil {
			// log scan errors to help debugging
			// return error instead of silently continuing
			return nil, err
		}
		if eventIDRaw.Valid {
			d.EventID = eventIDRaw.String
		} else {
			d.EventID = ""
		}
		if yearRaw.Valid {
			d.Year = int(yearRaw.Int64)
		} else {
			d.Year = 0
		}
		if imagesRaw.Valid && strings.TrimSpace(imagesRaw.String) != "" {
			var imgs []string
			if err := json.Unmarshal([]byte(imagesRaw.String), &imgs); err == nil {
				for i, im := range imgs {
					imgs[i] = normalizeImagePath(im)
				}
				d.Images = imgs
			} else {
				d.Images = []string{}
			}
		} else {
			d.Images = []string{}
		}
		res = append(res, d)
	}
	return res, nil
}

