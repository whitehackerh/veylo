package presenters

import (
	"time"

	"github.com/whitehackerh/veylo/api_server/src/presentation/openapi"
	"github.com/whitehackerh/veylo/api_server/src/usecases/dtos"
)

func ToLoginData(output *dtos.LoginOutput) openapi.LoginData {
	return openapi.LoginData{
		User: openapi.User{
			ID:        output.ID.String(),
			Email:     output.Email,
			CreatedAt: output.CreatedAt.UTC().Truncate(time.Second),
		},
		AccessToken: output.AccessToken,
		TokenType:   "Bearer",
	}
}
