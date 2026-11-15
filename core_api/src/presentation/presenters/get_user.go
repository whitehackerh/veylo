package presenters

import (
	"github.com/whitehackerh/veylo/api_server/src/generated/openapi"
	"github.com/whitehackerh/veylo/api_server/src/usecases/dtos"
)

type GetUserPresenter struct {
	userPresenter *UserPresenter
}

func NewGetUserPresenter(userPresenter *UserPresenter) *GetUserPresenter {
	return &GetUserPresenter{
		userPresenter: userPresenter,
	}
}

func (p *GetUserPresenter) ToResponse(output *dtos.GetUserOutput) openapi.User {
	return p.userPresenter.ToResponse(User{
		Id:        output.Id,
		Email:     output.Email,
		Name:      output.Name,
		CreatedAt: output.CreatedAt,
	})
}
