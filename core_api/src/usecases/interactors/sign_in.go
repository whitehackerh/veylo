package interactors

import (
	"context"
	"errors"
	"strings"

	"golang.org/x/crypto/bcrypt"

	"github.com/whitehackerh/veylo/api_server/src/usecases/dtos"
	"github.com/whitehackerh/veylo/api_server/src/usecases/repositories"
	"github.com/whitehackerh/veylo/api_server/src/usecases/security"
	"github.com/whitehackerh/veylo/api_server/src/usecases/shared"
)

type SignInInteractor struct {
	userRepository repositories.UserRepository
	tokenManager   security.TokenManager
}

func NewSignInInteractor(userRepository repositories.UserRepository, tokenManager security.TokenManager) *SignInInteractor {
	return &SignInInteractor{userRepository: userRepository, tokenManager: tokenManager}
}

func (i *SignInInteractor) Execute(ctx context.Context, input dtos.SignInInput) (dtos.SignInOutput, error) {
	email := strings.ToLower(strings.TrimSpace(input.Email))
	user, err := i.userRepository.FindByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, shared.ErrUserNotFound) {
			return dtos.SignInOutput{}, shared.ErrInvalidCredentials
		}
		return dtos.SignInOutput{}, err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(input.Password)); err != nil {
		return dtos.SignInOutput{}, shared.ErrInvalidCredentials
	}

	accessToken, err := i.tokenManager.Issue(user.Id.String())
	if err != nil {
		return dtos.SignInOutput{}, err
	}

	return dtos.SignInOutput{
		Id:          user.Id,
		Email:       user.Email,
		CreatedAt:   user.CreatedAt,
		AccessToken: accessToken,
	}, nil
}
