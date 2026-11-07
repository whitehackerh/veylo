package dtos

import "github.com/whitehackerh/veylo/api_server/src/domain"

type RegisterUserInput struct {
	Email    string
	Password string
}

type RegisterUserOutput struct {
	User *domain.User
}
