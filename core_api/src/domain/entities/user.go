package entities

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id           uuid.UUID
	Email        string
	PasswordHash string
	Name         string
	CreatedAt    time.Time
}

func NewUser(email string, passwordHash string, name string) *User {
	return &User{
		Id:           uuid.New(),
		Email:        email,
		PasswordHash: passwordHash,
		Name:         name,
		CreatedAt:    time.Now(),
	}
}

func RestoreUser(id uuid.UUID, email string, passwordHash string, name string, createdAt time.Time) *User {
	return &User{
		Id:           id,
		Email:        email,
		PasswordHash: passwordHash,
		Name:         name,
		CreatedAt:    createdAt,
	}
}
