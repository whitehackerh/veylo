package dtos

import (
	"time"

	"github.com/google/uuid"
)

type SignUpInput struct {
	Email    string
	Password string
}

type SignUpOutput struct {
	Id        uuid.UUID
	Email     string
	CreatedAt time.Time
}
