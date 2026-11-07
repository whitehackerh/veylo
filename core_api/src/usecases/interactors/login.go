package interactors

import (
	"context"
	"strings"

	"golang.org/x/crypto/bcrypt"

	"github.com/whitehackerh/veylo/api_server/src/usecases/dtos"
	"github.com/whitehackerh/veylo/api_server/src/usecases/ports"
	"github.com/whitehackerh/veylo/api_server/src/usecases/repositories"
	"github.com/whitehackerh/veylo/api_server/src/usecases/security"
	"github.com/whitehackerh/veylo/api_server/src/usecases/shared"
)

type LoginInteractor struct {
	userRepository repositories.UserRepository
	tokenManager   security.TokenManager
}

func NewLoginInteractor(userRepository repositories.UserRepository, tokenManager security.TokenManager) *LoginInteractor {
	return &LoginInteractor{userRepository: userRepository, tokenManager: tokenManager}
}

func (i *LoginInteractor) Execute(ctx context.Context, input dtos.LoginInput) (*dtos.LoginOutput, error) {
	email := strings.ToLower(strings.TrimSpace(input.Email))
	user, err := i.userRepository.FindByEmail(ctx, email)
	if err != nil {
		return nil, shared.ErrInvalidCredentials
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(input.Password)); err != nil {
		return nil, shared.ErrInvalidCredentials
	}

	accessToken, err := i.tokenManager.Issue(user.ID)
	if err != nil {
		return nil, err
	}

	return &dtos.LoginOutput{User: user, AccessToken: accessToken}, nil
}

var _ ports.LoginUseCase = (*LoginInteractor)(nil)
