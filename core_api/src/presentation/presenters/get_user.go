package presenters

import (
	"time"

	"github.com/whitehackerh/veylo/api_server/src/presentation/openapi"
	"github.com/whitehackerh/veylo/api_server/src/usecases/dtos"
)

func ToGetUserResponse(output *dtos.GetUserOutput) openapi.User {
	return openapi.User{
		ID: output.ID,
		Email: output.Email,
		CreatedAt: output.CreatedAt.UTC().Truncate(time.Second),
	}
}
