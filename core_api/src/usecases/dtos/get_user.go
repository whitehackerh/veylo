package dtos

import "github.com/whitehackerh/veylo/api_server/src/domain/entities"

type GetUserInput struct {
	UserID string
}

type GetUserOutput struct {
	User *entities.User
}
