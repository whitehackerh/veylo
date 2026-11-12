package interactors

import (
	"context"

	"github.com/google/uuid"

	"github.com/whitehackerh/veylo/api_server/src/usecases/dtos"
	"github.com/whitehackerh/veylo/api_server/src/usecases/repositories"
	"github.com/whitehackerh/veylo/api_server/src/usecases/shared"
)

type GetUserInteractor struct {
	userRepository repositories.UserRepository
}

func NewGetUserInteractor(userRepository repositories.UserRepository) *GetUserInteractor {
	return &GetUserInteractor{userRepository: userRepository}
}

func (i *GetUserInteractor) Execute(ctx context.Context, input dtos.GetUserInput) (dtos.GetUserOutput, error) {
	userId, err := uuid.Parse(input.UserId)
	if err != nil {
		return dtos.GetUserOutput{}, shared.ErrUserNotFound
	}

	user, err := i.userRepository.FindById(ctx, userId)
	if err != nil {
		return dtos.GetUserOutput{}, err
	}
	return dtos.GetUserOutput{
		Id:        user.Id,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}, nil
}
