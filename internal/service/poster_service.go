package service

import (
	"io"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/google/uuid"

	"github.com/IndalAwalaikal/coconut-event-hub/backend/internal/model"
	"github.com/IndalAwalaikal/coconut-event-hub/backend/internal/repository"
)

type PosterService interface {
    CreateFromForm(title, ptype, date string, file multipart.File, header *multipart.FileHeader) (*model.Poster, error)
    UpdateFromForm(id, title, ptype, date string, file multipart.File, header *multipart.FileHeader) (*model.Poster, error)
    Delete(id string) error
    List() ([]model.Poster, error)
    GetByID(id string) (*model.Poster, error)
}

type posterService struct {
    repo repository.PosterRepository
    storageDir string
}

func NewPosterService(r repository.PosterRepository, storageDir string) PosterService {
    return &posterService{repo: r, storageDir: storageDir}
}

func (s *posterService) List() ([]model.Poster, error) {
    return s.repo.List()
}

func (s *posterService) GetByID(id string) (*model.Poster, error) {
    return s.repo.GetByID(id)
}

func (s *posterService) CreateFromForm(title, ptype, date string, file multipart.File, header *multipart.FileHeader) (*model.Poster, error) {
    // save file
    var imgPath string
    if file != nil && header != nil {
        destDir := filepath.Join(s.storageDir, "posters")
        if err := os.MkdirAll(destDir, 0755); err != nil {
            return nil, err
        }
        // generate filename
        ext := filepath.Ext(header.Filename)
        filename := uuid.New().String() + ext
        dstPath := filepath.Join(destDir, filename)
        out, err := os.Create(dstPath)
        if err != nil {
            return nil, err
        }
        defer out.Close()
        if _, err := io.Copy(out, file); err != nil {
            return nil, err
        }
        // store web path
        imgPath = "/storage/posters/" + filename
    }

    p := &model.Poster{
        ID: uuid.New().String(),
        Title: title,
        Type: ptype,
        Image: imgPath,
        Date: date,
    }
    if err := s.repo.Create(p); err != nil {
        // cleanup file if persisted
        if imgPath != "" {
            _ = os.Remove(filepath.Join(s.storageDir, "posters", filepath.Base(imgPath)))
        }
        return nil, err
    }
    return p, nil
}

func (s *posterService) UpdateFromForm(id, title, ptype, date string, file multipart.File, header *multipart.FileHeader) (*model.Poster, error) {
    existing, err := s.repo.GetByID(id)
    if err != nil {
        return nil, err
    }
    var imgPath = existing.Image
    if file != nil && header != nil {
        destDir := filepath.Join(s.storageDir, "posters")
        if err := os.MkdirAll(destDir, 0755); err != nil {
            return nil, err
        }
        ext := filepath.Ext(header.Filename)
        filename := uuid.New().String() + ext
        dstPath := filepath.Join(destDir, filename)
        out, err := os.Create(dstPath)
        if err != nil {
            return nil, err
        }
        defer out.Close()
        if _, err := io.Copy(out, file); err != nil {
            return nil, err
        }
        // remove old file if present
        if existing.Image != "" {
            _ = os.Remove(filepath.Join(s.storageDir, "posters", filepath.Base(existing.Image)))
        }
        imgPath = "/storage/posters/" + filename
    }
    existing.Title = title
    existing.Type = ptype
    existing.Date = date
    existing.Image = imgPath
    if err := s.repo.Update(existing); err != nil {
        return nil, err
    }
    return existing, nil
}

func (s *posterService) Delete(id string) error {
    p, err := s.repo.GetByID(id)
    if err != nil {
        return err
    }
    if p.Image != "" {
        _ = os.Remove(filepath.Join(s.storageDir, "posters", filepath.Base(p.Image)))
    }
    return s.repo.Delete(id)
}
