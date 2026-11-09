package dtos

import (
	"time"

	"github.com/google/uuid"
)

type GetUserInput struct {
	UserID string
}

type GetUserOutput struct {
	ID        uuid.UUID
	Email     string
	CreatedAt time.Time
}
