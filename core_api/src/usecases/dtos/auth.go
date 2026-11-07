package dtos

import "github.com/whitehackerh/veylo/api_server/src/entities"

type RegisterUserInput struct {
	Email    string
	Password string
}

type LoginInput struct {
	Email    string
	Password string
}

type LoginOutput struct {
	User        *entities.User
	AccessToken string
}
