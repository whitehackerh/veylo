package interactors

import (
	"context"
	"net/mail"
	"strings"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

	"github.com/whitehackerh/veylo/api_server/src/domain/entities"
	"github.com/whitehackerh/veylo/api_server/src/usecases/dtos"
	"github.com/whitehackerh/veylo/api_server/src/usecases/repositories"
	"github.com/whitehackerh/veylo/api_server/src/usecases/shared"
)

type RegisterUserInteractor struct {
	userRepository repositories.UserRepository
}

func NewRegisterUserInteractor(userRepository repositories.UserRepository) *RegisterUserInteractor {
	return &RegisterUserInteractor{userRepository: userRepository}
}

func (i *RegisterUserInteractor) Execute(ctx context.Context, input dtos.RegisterUserInput) (*dtos.RegisterUserOutput, error) {
	email := strings.ToLower(strings.TrimSpace(input.Email))
	parsedEmail, err := mail.ParseAddress(email)
	if err != nil || parsedEmail.Address != email || len(input.Password) < 8 || len(input.Password) > 72 {
		return nil, shared.ErrInvalidRegistrationInput
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &entities.User{ID: uuid.NewString(), Email: email, PasswordHash: string(passwordHash)}
	if err := i.userRepository.Create(ctx, user); err != nil {
		return nil, err
	}

	return &dtos.RegisterUserOutput{User: user}, nil
}
