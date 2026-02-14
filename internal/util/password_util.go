package util

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
    bs, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return "", err
    }
    return string(bs), nil
}

func CompareHashAndPassword(hash, password string) error {
    return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
