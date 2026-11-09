package presenters

import (
	"github.com/whitehackerh/veylo/api_server/src/presentation/openapi"
	"github.com/whitehackerh/veylo/api_server/src/usecases/dtos"
)

func ToGetUserResponse(output *dtos.GetUserOutput) openapi.User {
	return ToUser(output)
}
