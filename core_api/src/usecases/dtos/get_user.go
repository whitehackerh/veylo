package dtos

import "github.com/whitehackerh/veylo/api_server/src/domain"

type GetUserInput struct {
	UserID string
}

type GetUserOutput struct {
	User *domain.User
}
