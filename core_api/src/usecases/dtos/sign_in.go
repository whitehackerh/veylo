package dtos

import (
	"time"

	"github.com/google/uuid"
)

type SignInInput struct {
	Email    string
	Password string
}

type SignInOutput struct {
	Id          uuid.UUID
	Email       string
	CreatedAt   time.Time
	AccessToken string
}
