package presenters

import (
	"time"

	"github.com/whitehackerh/veylo/api_server/src/domain/entities"
	"github.com/whitehackerh/veylo/api_server/src/presentation/openapi"
)

func ToGetUserResponse(user *entities.User) openapi.User {
	return openapi.User{
		ID: user.ID,
		Email: user.Email,
		CreatedAt: user.CreatedAt.UTC().Truncate(time.Second),
	}
}
