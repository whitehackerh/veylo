package presenters

import (
	"time"

	"github.com/google/uuid"
	openapi_types "github.com/oapi-codegen/runtime/types"

	"github.com/whitehackerh/veylo/api_server/src/generated/openapi"
)

type User struct {
	Id        uuid.UUID
	Email     string
	CreatedAt time.Time
}

func ToUser(user User) openapi.User {
	return openapi.User{
		Id:    openapi_types.UUID(user.Id),
		Email: openapi_types.Email(user.Email),
		CreatedAt: user.CreatedAt.UTC().Truncate(time.Second),
	}
}
