package presenters

import (
	"time"

	"github.com/whitehackerh/veylo/api_server/src/presentation/openapi"
	"github.com/whitehackerh/veylo/api_server/src/usecases/dtos"
)

func ToGetUserResponse(output *dtos.GetUserOutput, path string) openapi.GetUserResponse {
	return openapi.GetUserResponse{
		Data: openapi.User{
			ID:        output.ID.String(),
			Email:     output.Email,
			CreatedAt: output.CreatedAt.UTC().Truncate(time.Second),
		},
		CommonResponse: ToCommonResponse(path),
	}
}
