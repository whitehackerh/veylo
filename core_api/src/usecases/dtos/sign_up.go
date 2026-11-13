package dtos

import (
	"time"

	"github.com/google/uuid"
)

type SignUpInput struct {
	Email    string
	Password string
	Name     string
}

type SignUpOutput struct {
	Id        uuid.UUID
	Email     string
	Name      string
	CreatedAt time.Time
}
