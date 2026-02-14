package service

import (
	"database/sql"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"

	"github.com/google/uuid"

	"github.com/IndalAwalaikal/coconut-event-hub/backend/internal/model"
	"github.com/IndalAwalaikal/coconut-event-hub/backend/internal/repository"
)

type RegistrationService struct {
    DB *sql.DB
    Repo *repository.RegistrationRepository
    StoragePath string // directory to store uploaded proof files
}

func NewRegistrationService(db *sql.DB, repo *repository.RegistrationRepository, storage string) *RegistrationService {
    return &RegistrationService{DB: db, Repo: repo, StoragePath: storage}
}

// saveUploadedFile writes the uploaded multipart file to storage and returns relative path
func (s *RegistrationService) saveUploadedFile(file multipart.File, header *multipart.FileHeader) (string, error) {
    if s.StoragePath == "" {
        s.StoragePath = "storage/registrations"
    }
    if err := os.MkdirAll(s.StoragePath, 0o755); err != nil {
        return "", err
    }
    ext := filepath.Ext(header.Filename)
    fname := fmt.Sprintf("%s%s", uuid.New().String(), ext)
    dest := filepath.Join(s.StoragePath, fname)
    out, err := os.Create(dest)
    if err != nil {
        return "", err
    }
    defer out.Close()
    if _, err := io.Copy(out, file); err != nil {
        return "", err
    }
    // return path relative to project root
    return "/" + filepath.ToSlash(dest), nil
}

// Register handles registration creation with optional file upload.
func (s *RegistrationService) Register(eventID, name, whatsapp, institution string, file multipart.File, header *multipart.FileHeader, fileName string) (*model.Registration, error) {
    if eventID == "" || name == "" || whatsapp == "" {
        return nil, errors.New("missing required fields")
    }

    tx, err := s.DB.Begin()
    if err != nil {
        return nil, err
    }
    defer func() {
        // rollback if not committed
        _ = tx.Rollback()
    }()

    // lock the event row to safely check and update quota
    var registered, quota int
    q := `SELECT registered, quota FROM events WHERE id = ? FOR UPDATE`
    if err := tx.QueryRow(q, eventID).Scan(&registered, &quota); err != nil {
        if err == sql.ErrNoRows {
            return nil, errors.New("event not found")
        }
        return nil, err
    }
    if registered >= quota {
        return nil, errors.New("quota full")
    }

    // handle file upload if provided
    var proofPath string
    if file != nil && header != nil {
        p, err := s.saveUploadedFile(file, header)
        if err != nil {
            return nil, err
        }
        proofPath = p
        if fileName == "" {
            fileName = header.Filename
        }
    }

    reg := &model.Registration{
        ID: uuid.New().String(),
        EventID: eventID,
        Name: name,
        Whatsapp: whatsapp,
        Institution: institution,
        ProofImage: proofPath,
        FileName: fileName,
        RegisteredAt: time.Now(),
    }

    // insert registration using tx
    ins := `INSERT INTO registrations (id,event_id,name,whatsapp,institution,proof_image,file_name,registered_at,created_at,updated_at) VALUES (?,?,?,?,?,?,?,?,NOW(),NOW())`
    if _, err := tx.Exec(ins, reg.ID, reg.EventID, reg.Name, reg.Whatsapp, reg.Institution, reg.ProofImage, reg.FileName, reg.RegisteredAt); err != nil {
        return nil, err
    }

    // increment registered count
    if _, err := tx.Exec(`UPDATE events SET registered = registered + 1 WHERE id = ?`, eventID); err != nil {
        return nil, err
    }

    if err := tx.Commit(); err != nil {
        return nil, err
    }

    return reg, nil
}

// ListRegistrationsByEvent returns a slice of registrations for an event
func (s *RegistrationService) ListRegistrationsByEvent(eventID string) ([]model.Registration, error) {
    return s.Repo.ListByEvent(eventID)
}

// GetRegistration returns a registration by id
func (s *RegistrationService) GetRegistration(id string) (*model.Registration, error) {
    return s.Repo.GetByID(id)
}
