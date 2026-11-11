package entities

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id           uuid.UUID
	Email        string
	PasswordHash string
	CreatedAt    time.Time
}

func NewUser(email string, passwordHash string) *User {
	return &User{
		Id:           uuid.New(),
		Email:        email,
		PasswordHash: passwordHash,
		CreatedAt:    time.Now(),
	}
}

func RestoreUser(id uuid.UUID, email string, passwordHash string, createdAt time.Time) *User {
	return &User{
		Id:           id,
		Email:        email,
		PasswordHash: passwordHash,
		CreatedAt:    createdAt,
	}
}
