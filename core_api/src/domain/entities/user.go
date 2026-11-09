package entities

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID           string
	Email        string
	PasswordHash string
	CreatedAt    time.Time
}

func NewUser(email string, passwordHash string) *User {
	return &User{
		ID:           uuid.NewString(),
		Email:        email,
		PasswordHash: passwordHash,
	}
}

func RestoreUser(id string, email string, passwordHash string, createdAt time.Time) *User {
	return &User{
		ID:           id,
		Email:        email,
		PasswordHash: passwordHash,
		CreatedAt:    createdAt,
	}
}
