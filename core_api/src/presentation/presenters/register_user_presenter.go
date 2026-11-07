package presenters

import (
	"time"

	"github.com/whitehackerh/veylo/api_server/src/usecases/dtos"
)

type RegisterUserResponse struct {
	ID        string `json:"id"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
}

func ToRegisterUserResponse(output *dtos.RegisterUserOutput) RegisterUserResponse {
	return RegisterUserResponse{
		ID: output.User.ID,
		Email: output.User.Email,
		CreatedAt: output.User.CreatedAt.UTC().Format(time.RFC3339),
	}
}
