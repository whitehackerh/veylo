package interactors

import (
	"context"

	"github.com/whitehackerh/veylo/api_server/src/usecases/dtos"
	"github.com/whitehackerh/veylo/api_server/src/usecases/repositories"
)

type GetUserInteractor struct {
	userRepository repositories.UserRepository
}

func NewGetUserInteractor(userRepository repositories.UserRepository) *GetUserInteractor {
	return &GetUserInteractor{userRepository: userRepository}
}

func (i *GetUserInteractor) Execute(ctx context.Context, input dtos.GetUserInput) (*dtos.GetUserOutput, error) {
	user, err := i.userRepository.FindByID(ctx, input.UserID)
	if err != nil {
		return nil, err
	}
	return &dtos.GetUserOutput{User: user}, nil
}
