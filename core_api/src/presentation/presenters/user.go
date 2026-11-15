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
	Name      string
	CreatedAt time.Time
}

type UserPresenter struct{}

func NewUserPresenter() *UserPresenter {
	return &UserPresenter{}
}

func (p *UserPresenter) ToResponse(user User) openapi.User {
	return openapi.User{
		CreatedAt: user.CreatedAt.UTC().Truncate(time.Second),
		Email:     openapi_types.Email(user.Email),
		Id:        openapi_types.UUID(user.Id),
		Name:      user.Name,
	}
}
