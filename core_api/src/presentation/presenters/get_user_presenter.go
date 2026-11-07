package presenters

import (
	"time"

	"github.com/whitehackerh/veylo/api_server/src/usecases/dtos"
)

type GetUserResponse struct {
	ID        string `json:"id"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
}

func ToGetUserResponse(output *dtos.GetUserOutput) GetUserResponse {
	return GetUserResponse{
		ID: output.User.ID,
		Email: output.User.Email,
		CreatedAt: output.User.CreatedAt.UTC().Format(time.RFC3339),
	}
}
