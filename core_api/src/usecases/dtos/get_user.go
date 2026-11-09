package dtos

import "time"

type GetUserInput struct {
	UserID string
}

type GetUserOutput struct {
	ID        string
	Email     string
	CreatedAt time.Time
}
