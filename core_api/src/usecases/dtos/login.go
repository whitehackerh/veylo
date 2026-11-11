package dtos

import (
	"time"

	"github.com/google/uuid"
)

type LoginInput struct {
	Email    string
	Password string
}

type LoginOutput struct {
	Id          uuid.UUID
	Email       string
	CreatedAt   time.Time
	AccessToken string
}
