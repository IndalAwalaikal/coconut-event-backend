package service

import (
	"database/sql"
	"errors"
	"time"

	"github.com/IndalAwalaikal/coconut-event-hub/backend/internal/repository"
	"github.com/IndalAwalaikal/coconut-event-hub/backend/internal/util"
	"github.com/google/uuid"
)

type AuthService struct {
    db *sql.DB
    adminRepo *repository.AdminRepository
}

func NewAuthService(db *sql.DB, adminRepo *repository.AdminRepository) *AuthService {
    return &AuthService{db: db, adminRepo: adminRepo}
}

func (s *AuthService) Login(username, password string) (string, error) {
    admin, err := s.adminRepo.GetByUsername(username)
    if err != nil {
        if err == sql.ErrNoRows { return "", errors.New("invalid credentials") }
        return "", err
    }
    if err := util.CompareHashAndPassword(admin.PasswordHash, password); err != nil {
        return "", errors.New("invalid credentials")
    }
    // generate token
    token, err := util.GenerateAdminToken(int64(admin.ID), admin.Username, time.Hour*24)
    if err != nil {
        return "", err
    }
    // optional: update last_login or token id
    _ = uuid.New()
    return token, nil
}
