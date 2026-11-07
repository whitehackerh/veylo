package interactors

import (
	"context"

	"github.com/whitehackerh/veylo/api_server/src/entities"
	"github.com/whitehackerh/veylo/api_server/src/usecases/repositories"
)

type GetUserInteractor struct {
	userRepository repositories.UserRepository
}

func NewGetUserInteractor(userRepository repositories.UserRepository) *GetUserInteractor {
	return &GetUserInteractor{userRepository: userRepository}
}

func (i *GetUserInteractor) Execute(ctx context.Context, userID string) (*entities.User, error) {
	return i.userRepository.FindByID(ctx, userID)
}
