package dtos

import "time"

type RegisterUserInput struct {
	Email    string
	Password string
}

type RegisterUserOutput struct {
	ID        string
	Email     string
	CreatedAt time.Time
}
