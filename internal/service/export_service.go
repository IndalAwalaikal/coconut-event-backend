package service

import (
	"bytes"
	"encoding/csv"
	"fmt"

	"github.com/IndalAwalaikal/coconut-event-hub/backend/internal/repository"
)

type ExportService struct {
    Repo *repository.RegistrationRepository
}

func NewExportService(repo *repository.RegistrationRepository) *ExportService {
    return &ExportService{Repo: repo}
}

// ExportRegistrationsCSV returns CSV bytes for registrations of an event
func (s *ExportService) ExportRegistrationsCSV(eventID string) ([]byte, error) {
    regs, err := s.Repo.ListByEvent(eventID)
    if err != nil {
        return nil, err
    }
    buf := &bytes.Buffer{}
    w := csv.NewWriter(buf)

    // header
    header := []string{"ID", "Name", "WhatsApp", "Institution", "FileName", "ProofImage", "RegisteredAt"}
    if err := w.Write(header); err != nil {
        return nil, err
    }

    for _, r := range regs {
        rec := []string{
            r.ID,
            r.Name,
            r.Whatsapp,
            r.Institution,
            r.FileName,
            r.ProofImage,
            r.RegisteredAt.Format("2006-01-02 15:04:05"),
        }
        if err := w.Write(rec); err != nil {
            return nil, fmt.Errorf("write csv: %w", err)
        }
    }
    w.Flush()
    if err := w.Error(); err != nil {
        return nil, err
    }
    return buf.Bytes(), nil
}
