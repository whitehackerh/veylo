package dtos

import (
	"time"

	"github.com/google/uuid"
)

type RegisterUserInput struct {
	Email    string
	Password string
}

type RegisterUserOutput struct {
	ID        uuid.UUID
	Email     string
	CreatedAt time.Time
}
