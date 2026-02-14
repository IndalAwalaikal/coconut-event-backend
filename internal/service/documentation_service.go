package service

import (
	"database/sql"
	"errors"
	"fmt"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/IndalAwalaikal/coconut-event-hub/backend/internal/model"
	"github.com/IndalAwalaikal/coconut-event-hub/backend/internal/repository"
	"github.com/IndalAwalaikal/coconut-event-hub/backend/internal/util"
	"github.com/google/uuid"
)

// documentation service
type DocumentationService struct {
	db   *sql.DB
	repo *repository.DocumentationRepository
	storageDir string
}

func NewDocumentationService(db *sql.DB, repo *repository.DocumentationRepository, storageDir string) *DocumentationService {
	return &DocumentationService{db: db, repo: repo, storageDir: storageDir}
}

// DebugDatabase returns the current database name for diagnostic purposes
func (s *DocumentationService) DebugDatabase() (string, error) {
    var dbname sql.NullString
    if err := s.db.QueryRow(`SELECT DATABASE()`).Scan(&dbname); err != nil {
        return "", err
    }
    if dbname.Valid {
        return dbname.String, nil
    }
    return "", nil
}

func (s *DocumentationService) List(category string, year int, q string, limit, offset int) ([]model.Documentation, error) {
	return s.repo.List(category, year, q, limit, offset)
}

func (s *DocumentationService) GetByID(id string) (*model.Documentation, error) {
	return s.repo.GetByID(id)
}

func (s *DocumentationService) CreateFromForm(r *http.Request) (*model.Documentation, error) {
    if err := r.ParseMultipartForm(32 << 20); err != nil {
        return nil, err
    }
    
    // parse fields
    category := r.FormValue("category")
    categoryLabel := r.FormValue("category_label")
    eventTitle := r.FormValue("event_title")
    yearStr := r.FormValue("year")
    year := 0
    if yearStr != "" {
        if y, err := strconv.Atoi(yearStr); err == nil { year = y }
    }
    description := r.FormValue("description")
    eventID := r.FormValue("event_id")

    // files
    var files []*multipart.FileHeader
    if r.MultipartForm != nil {
        files = r.MultipartForm.File["images"]
    }
    if len(files) == 0 {
        return nil, errors.New("at least one image required")
    }
    if len(files) > 10 {
        return nil, errors.New("maximum 10 images allowed")
    }
    
    var saved []string
    var cleanupFiles []string // track files to cleanup on error
    
    for _, fh := range files {
        f, err := fh.Open()
        if err != nil {
            // cleanup previously saved files
            for _, p := range cleanupFiles {
                os.Remove(strings.TrimPrefix(p, "/"))
            }
            return nil, fmt.Errorf("failed to open file %s: %w", fh.Filename, err)
        }
        
        destDir := filepath.Join(s.storageDir)
        p, err := util.SaveMultipartFile(destDir, f, fh)
        f.Close()
        
        if err != nil {
            // cleanup previously saved files
            for _, p := range cleanupFiles {
                os.Remove(strings.TrimPrefix(p, "/"))
            }
            return nil, fmt.Errorf("failed to save file %s: %w", fh.Filename, err)
        }
        
        saved = append(saved, p)
        cleanupFiles = append(cleanupFiles, p)
    }

    doc := &model.Documentation{
        ID:            uuid.New().String(),
        EventID:       eventID,
        Category:      category,
        CategoryLabel: categoryLabel,
        EventTitle:    eventTitle,
        Year:          year,
        Images:        saved,
        Description:   description,
        CreatedAt:     time.Now(),
        UpdatedAt:     time.Now(),
    }
    
    if err := s.repo.Create(doc); err != nil {
        // cleanup all saved files if DB insert fails
        for _, p := range saved {
            os.Remove(strings.TrimPrefix(p, "/"))
        }
        return nil, fmt.Errorf("failed to save to database: %w", err)
    }
    
    return doc, nil
}

func (s *DocumentationService) UpdateFromForm(id string, r *http.Request) (*model.Documentation, error) {
    existing, err := s.repo.GetByID(id)
    if err != nil { 
        return nil, fmt.Errorf("documentation not found: %w", err) 
    }
    
    if err := r.ParseMultipartForm(32 << 20); err != nil { 
        return nil, err 
    }
    
    // parse fields (optional)
    if v := r.FormValue("category"); v != "" { existing.Category = v }
    if v := r.FormValue("category_label"); v != "" { existing.CategoryLabel = v }
    if v := r.FormValue("event_title"); v != "" { existing.EventTitle = v }
    if v := r.FormValue("year"); v != "" { 
        if y, e := strconv.Atoi(v); e == nil { existing.Year = y } 
    }
    if v := r.FormValue("description"); v != "" { existing.Description = v }

    // optional additional images (append) - enforce max 10 total
    var files []*multipart.FileHeader
    if r.MultipartForm != nil {
        files = r.MultipartForm.File["images"]
    }
    
    var cleanupFiles []string
    var newImages []string

    if len(files) > 0 {
        if len(existing.Images)+len(files) > 10 {
            return nil, errors.New("maximum 10 images allowed")
        }

        for _, fh := range files {
            f, err := fh.Open()
            if err != nil {
                // cleanup any newly saved files
                for _, p := range cleanupFiles {
                    os.Remove(strings.TrimPrefix(p, "/"))
                }
                return nil, fmt.Errorf("failed to open file %s: %w", fh.Filename, err)
            }
            
            p, err := util.SaveMultipartFile(s.storageDir, f, fh)
            f.Close()
            
            if err != nil {
                // cleanup any newly saved files
                for _, p := range cleanupFiles {
                    os.Remove(strings.TrimPrefix(p, "/"))
                }
                return nil, fmt.Errorf("failed to save file %s: %w", fh.Filename, err)
            }
            
            newImages = append(newImages, p)
            cleanupFiles = append(cleanupFiles, p)
        }
        
        // append new images to existing
        existing.Images = append(existing.Images, newImages...)
    }
    
    existing.UpdatedAt = time.Now()
    
    if err := s.repo.Update(existing); err != nil {
        // cleanup newly saved files if update fails
        for _, p := range cleanupFiles {
            os.Remove(strings.TrimPrefix(p, "/"))
        }
        return nil, fmt.Errorf("failed to update database: %w", err)
    }
    
    return existing, nil
}

func (s *DocumentationService) Delete(id string) error {
	return s.repo.Delete(id)
}

