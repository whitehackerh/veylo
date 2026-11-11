package presenters

import (
	"github.com/whitehackerh/veylo/api_server/src/generated/openapi"
	"github.com/whitehackerh/veylo/api_server/src/usecases/dtos"
)

func ToGetUserResponse(output *dtos.GetUserOutput) openapi.User {
	return ToUser(User{
		Id: output.Id,
		Email: output.Email,
		CreatedAt: output.CreatedAt,
	})
}
