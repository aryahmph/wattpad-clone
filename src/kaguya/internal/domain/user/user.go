package user

import (
	"time"
)

type User struct {
	ID           string    `json:"id"`
	Username     string    `json:"username"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"password_hash"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type Register struct {
	Username string `validate:"required,min=5,max=100"`
	Email    string `validate:"required,email,min=5,max=255"`
	Password string `validate:"required,min=8,max=16"`
}

type Login struct {
	Username string `validate:"required,min=5,max=100"`
	Password string `validate:"required,min=8,max=16"`
}
