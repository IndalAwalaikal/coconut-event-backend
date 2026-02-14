package model

import "time"

type Admin struct {
    ID          int64     `db:"id" json:"id"`
    Username    string    `db:"username" json:"username"`
    PasswordHash string   `db:"password_hash" json:"-"`
    Name        string    `db:"name" json:"name"`
    Role        string    `db:"role" json:"role"`
    CreatedAt   time.Time `db:"created_at" json:"createdAt"`
    UpdatedAt   time.Time `db:"updated_at" json:"updatedAt"`
}
