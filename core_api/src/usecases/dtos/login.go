package dtos

import "github.com/whitehackerh/veylo/api_server/src/domain"

type LoginInput struct {
	Email    string
	Password string
}

type LoginOutput struct {
	User        *domain.User
	AccessToken string
}
