package util

import (
	"errors"
	"os"
	"time"

	jwt "github.com/golang-jwt/jwt/v4"
)

type AdminClaims struct {
    AdminID int64  `json:"adminId"`
    Username string `json:"username"`
    jwt.RegisteredClaims
}

func jwtSecret() []byte {
    s := os.Getenv("JWT_SECRET")
    if s == "" {
        s = "change-me-in-production"
    }
    return []byte(s)
}

func GenerateAdminToken(adminID int64, username string, ttl time.Duration) (string, error) {
    if ttl == 0 { ttl = time.Hour * 24 }
    now := time.Now()
    claims := AdminClaims{
        AdminID: adminID,
        Username: username,
        RegisteredClaims: jwt.RegisteredClaims{
            Subject: username,
            IssuedAt: jwt.NewNumericDate(now),
            ExpiresAt: jwt.NewNumericDate(now.Add(ttl)),
        },
    }
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(jwtSecret())
}

func ParseAdminToken(tokenStr string) (*AdminClaims, error) {
    token, err := jwt.ParseWithClaims(tokenStr, &AdminClaims{}, func(t *jwt.Token) (interface{}, error) { return jwtSecret(), nil })
    if err != nil {
        return nil, err
    }
    if !token.Valid {
        return nil, errors.New("invalid token")
    }
    if claims, ok := token.Claims.(*AdminClaims); ok {
        return claims, nil
    }
    return nil, errors.New("invalid claims")
}
