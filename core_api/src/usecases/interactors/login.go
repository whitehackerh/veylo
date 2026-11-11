package interactors

import (
	"context"
	"strings"

	"golang.org/x/crypto/bcrypt"

	"github.com/whitehackerh/veylo/api_server/src/usecases/dtos"
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

func (i *LoginInteractor) Execute(ctx context.Context, input dtos.LoginInput) (dtos.LoginOutput, error) {
	email := strings.ToLower(strings.TrimSpace(input.Email))
	user, err := i.userRepository.FindByEmail(ctx, email)
	if err != nil {
		return dtos.LoginOutput{}, shared.ErrInvalidCredentials
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(input.Password)); err != nil {
		return dtos.LoginOutput{}, shared.ErrInvalidCredentials
	}

	accessToken, err := i.tokenManager.Issue(user.Id.String())
	if err != nil {
		return dtos.LoginOutput{}, err
	}

	return dtos.LoginOutput{
		Id:          user.Id,
		Email:       user.Email,
		CreatedAt:   user.CreatedAt,
		AccessToken: accessToken,
	}, nil
}
