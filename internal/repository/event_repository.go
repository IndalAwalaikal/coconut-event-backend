package repository

import (
	"database/sql"
	"encoding/json"
	"log"
	"time"

	"github.com/IndalAwalaikal/coconut-event-hub/backend/internal/model"
)

type EventRepository struct {
    DB *sql.DB
}

func NewEventRepository(db *sql.DB) *EventRepository {
    return &EventRepository{DB: db}
}

func (r *EventRepository) Create(e *model.Event) error {
    rulesJson, _ := json.Marshal(e.Rules)
    benefitsJson, _ := json.Marshal(e.Benefits)
    var date interface{}
    if e.Date != nil {
        date = e.Date.Format("2006-01-02")
    } else {
        date = nil
    }
    q := `INSERT INTO events (id,category,category_label,title,description,rules,benefits,date,time,speaker1,speaker2,speaker3,moderator,location,quota,registered,poster,event_type,price,available,created_at,updated_at) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,NOW(),NOW())`
    _, err := r.DB.Exec(q, e.ID, e.Category, e.CategoryLabel, e.Title, e.Description, string(rulesJson), string(benefitsJson), date, e.Time, e.Speaker1, e.Speaker2, e.Speaker3, e.Moderator, e.Location, e.Quota, e.Registered, e.Poster, e.EventType, e.Price, e.Available)
    return err
}

func (r *EventRepository) Update(e *model.Event) error {
    rulesJson, _ := json.Marshal(e.Rules)
    benefitsJson, _ := json.Marshal(e.Benefits)
    var date interface{}
    if e.Date != nil {
        date = e.Date.Format("2006-01-02")
    } else {
        date = nil
    }
    q := `UPDATE events SET category=?,category_label=?,title=?,description=?,rules=?,benefits=?,date=?,time=?,speaker1=?,speaker2=?,speaker3=?,moderator=?,location=?,quota=?,poster=?,event_type=?,price=?,available=?,updated_at=NOW() WHERE id=?`
    _, err := r.DB.Exec(q, e.Category, e.CategoryLabel, e.Title, e.Description, string(rulesJson), string(benefitsJson), date, e.Time, e.Speaker1, e.Speaker2, e.Speaker3, e.Moderator, e.Location, e.Quota, e.Poster, e.EventType, e.Price, e.Available, e.ID)
    return err
}

func (r *EventRepository) Delete(id string) error {
    _, err := r.DB.Exec(`DELETE FROM events WHERE id = ?`, id)
    return err
}

func (r *EventRepository) GetByID(id string) (*model.Event, error) {
    q := `SELECT id,category,category_label,title,description,rules,benefits,date,time,speaker1,speaker2,speaker3,moderator,location,quota,registered,poster,event_type,price,available,created_at,updated_at FROM events WHERE id = ? LIMIT 1`
    row := r.DB.QueryRow(q, id)
    var e model.Event
    var rulesStr, benefitsStr sql.NullString
    var date sql.NullTime
    var timeStr, speaker1, speaker2, speaker3, moderator, location sql.NullString
    var eventType sql.NullString
    var price int
    var available int
    if err := row.Scan(&e.ID, &e.Category, &e.CategoryLabel, &e.Title, &e.Description, &rulesStr, &benefitsStr, &date, &timeStr, &speaker1, &speaker2, &speaker3, &moderator, &location, &e.Quota, &e.Registered, &e.Poster, &eventType, &price, &available, &e.CreatedAt, &e.UpdatedAt); err != nil {
        return nil, err
    }
    if rulesStr.Valid {
        _ = json.Unmarshal([]byte(rulesStr.String), &e.Rules)
        if e.Rules == nil {
            e.Rules = []string{}
        }
    } else {
        e.Rules = []string{}
    }
    if benefitsStr.Valid {
        _ = json.Unmarshal([]byte(benefitsStr.String), &e.Benefits)
        if e.Benefits == nil {
            e.Benefits = []string{}
        }
    } else {
        e.Benefits = []string{}
    }
    if date.Valid {
        t := date.Time
        e.Date = &t
        y1, m1, d1 := t.Date()
        y2, m2, d2 := time.Now().Date()
        if y1 < y2 || (y1 == y2 && (m1 < m2 || (m1 == m2 && d1 < d2))) {
            e.Available = false
        } else {
            e.Available = available != 0
        }
    }
    if timeStr.Valid { e.Time = timeStr.String }
    if speaker1.Valid { e.Speaker1 = speaker1.String }
    if speaker2.Valid { e.Speaker2 = speaker2.String }
    if speaker3.Valid { e.Speaker3 = speaker3.String }
    if moderator.Valid { e.Moderator = moderator.String }
    if location.Valid { e.Location = location.String }
    if eventType.Valid { e.EventType = eventType.String }
    e.Price = price
    e.Available = available != 0
    return &e, nil
}

// List with optional category and search (simple implementation)
func (r *EventRepository) List(category, search string) ([]model.Event, error) {
    log.Printf("EventRepository.List called - category: %q, search: %q", category, search)
    
    base := `SELECT id,category,category_label,title,description,rules,benefits,date,time,speaker1,speaker2,speaker3,moderator,location,quota,registered,poster,event_type,price,available,created_at,updated_at FROM events`
    var rows *sql.Rows
    var err error
    
    if category != "" && search != "" {
        q := base + ` WHERE category = ? AND (title LIKE ? OR description LIKE ?) ORDER BY COALESCE(date, created_at) DESC, created_at DESC`
        rows, err = r.DB.Query(q, category, "%"+search+"%", "%"+search+"%")
    } else if category != "" {
        q := base + ` WHERE category = ? ORDER BY COALESCE(date, created_at) DESC, created_at DESC`
        rows, err = r.DB.Query(q, category)
    } else if search != "" {
        q := base + ` WHERE title LIKE ? OR description LIKE ? ORDER BY COALESCE(date, created_at) DESC, created_at DESC`
        rows, err = r.DB.Query(q, "%"+search+"%", "%"+search+"%")
    } else {
        q := base + ` ORDER BY COALESCE(date, created_at) DESC, created_at DESC`
        rows, err = r.DB.Query(q)
    }
    
    if err != nil {
        log.Printf("Query error: %v", err)
        return nil, err
    }
    defer rows.Close()
    
    out := []model.Event{}
    count := 0
    
    for rows.Next() {
        count++
        var e model.Event
        var rulesStr, benefitsStr sql.NullString
        var date sql.NullTime
        var timeStr, speaker1, speaker2, speaker3, moderator, location sql.NullString
        var eventType sql.NullString
        var price int
        var available int
        
        err := rows.Scan(
            &e.ID, &e.Category, &e.CategoryLabel, &e.Title, &e.Description,
            &rulesStr, &benefitsStr, &date, &timeStr,
            &speaker1, &speaker2, &speaker3, &moderator, &location,
            &e.Quota, &e.Registered, &e.Poster, &eventType, &price, &available,
            &e.CreatedAt, &e.UpdatedAt,
        )
        
        if err != nil {
            log.Printf("Row scan error (row %d): %v", count, err)
            return nil, err
        }
        
        if rulesStr.Valid {
            if err := json.Unmarshal([]byte(rulesStr.String), &e.Rules); err != nil {
                log.Printf("Failed to unmarshal rules for event %s: %v", e.ID, err)
            }
            if e.Rules == nil { e.Rules = []string{} }
        } else {
            e.Rules = []string{}
        }

        if benefitsStr.Valid {
            if err := json.Unmarshal([]byte(benefitsStr.String), &e.Benefits); err != nil {
                log.Printf("Failed to unmarshal benefits for event %s: %v", e.ID, err)
            }
            if e.Benefits == nil { e.Benefits = []string{} }
        } else {
            e.Benefits = []string{}
        }
        
        if date.Valid {
            t := date.Time
            e.Date = &t
            y1, m1, d1 := t.Date()
            y2, m2, d2 := time.Now().Date()
            if y1 < y2 || (y1 == y2 && (m1 < m2 || (m1 == m2 && d1 < d2))) {
                e.Available = false
            } else {
                e.Available = available != 0
            }
        }
        
        if timeStr.Valid {
            e.Time = timeStr.String
        }
        
        if speaker1.Valid {
            e.Speaker1 = speaker1.String
        }
        
        if speaker2.Valid {
            e.Speaker2 = speaker2.String
        }
        
        if speaker3.Valid {
            e.Speaker3 = speaker3.String
        }
        
        if moderator.Valid {
            e.Moderator = moderator.String
        }
        
        if location.Valid {
            e.Location = location.String
        }
        
        if eventType.Valid {
            e.EventType = eventType.String
        }
        
        e.Price = price
        e.Available = available != 0
        
        out = append(out, e)
    }
    
    if err := rows.Err(); err != nil {
        log.Printf("Rows iteration error: %v", err)
        return nil, err
    }
    
    log.Printf("Successfully fetched %d events", count)
    
    return out, nil
}