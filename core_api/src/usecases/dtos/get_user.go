package dtos

import (
	"time"

	"github.com/google/uuid"
)

type GetUserInput struct {
	UserId string
}

type GetUserOutput struct {
	Id        uuid.UUID
	Email     string
	CreatedAt time.Time
}
