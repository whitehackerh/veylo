package presenters

import (
	"time"

	"github.com/whitehackerh/veylo/api_server/src/domain/entities"
)

type UserResponse struct {
	ID        string `json:"id"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
}

type LoginResponse struct {
	User        UserResponse `json:"user"`
	AccessToken string       `json:"access_token"`
	TokenType   string       `json:"token_type"`
}

func ToUserResponse(user *entities.User) UserResponse {
	return UserResponse{
		ID: user.ID, Email: user.Email,
		CreatedAt: user.CreatedAt.UTC().Format(time.RFC3339),
	}
}
