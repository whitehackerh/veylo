package presenters

import (
	"time"

	"github.com/whitehackerh/veylo/api_server/src/usecases/dtos"
)

type LoginUserResponse struct {
	ID        string `json:"id"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
}

type LoginResponse struct {
	User        LoginUserResponse `json:"user"`
	AccessToken string            `json:"access_token"`
	TokenType   string            `json:"token_type"`
}

func ToLoginResponse(output *dtos.LoginOutput) LoginResponse {
	return LoginResponse{
		User: LoginUserResponse{
			ID: output.User.ID,
			Email: output.User.Email,
			CreatedAt: output.User.CreatedAt.UTC().Format(time.RFC3339),
		},
		AccessToken: output.AccessToken,
		TokenType: "Bearer",
	}
}
