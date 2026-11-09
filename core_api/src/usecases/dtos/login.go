package dtos

import "time"

type LoginInput struct {
	Email    string
	Password string
}

type LoginOutput struct {
	ID          string
	Email       string
	CreatedAt   time.Time
	AccessToken string
}
