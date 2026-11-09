package presenters

import (
	"time"

	"github.com/whitehackerh/veylo/api_server/src/presentation/openapi"
	"github.com/whitehackerh/veylo/api_server/src/usecases/dtos"
)

func ToLoginResponse(output *dtos.LoginOutput) openapi.LoginResponse {
	return openapi.LoginResponse{
		User: openapi.User{
			ID:        output.ID.String(),
			Email:     output.Email,
			CreatedAt: output.CreatedAt.UTC().Truncate(time.Second),
		},
		AccessToken: output.AccessToken,
		TokenType:   "Bearer",
	}
}
