package interactors

import (
	"context"
	"net/mail"
	"strings"

	"golang.org/x/crypto/bcrypt"

	"github.com/whitehackerh/veylo/api_server/src/domain/entities"
	"github.com/whitehackerh/veylo/api_server/src/usecases/dtos"
	"github.com/whitehackerh/veylo/api_server/src/usecases/repositories"
	"github.com/whitehackerh/veylo/api_server/src/usecases/shared"
)

type SignUpInteractor struct {
	userRepository repositories.UserRepository
}

func NewSignUpInteractor(userRepository repositories.UserRepository) *SignUpInteractor {
	return &SignUpInteractor{userRepository: userRepository}
}

func (i *SignUpInteractor) Execute(ctx context.Context, input dtos.SignUpInput) (dtos.SignUpOutput, error) {
	if input.Email == "" || input.Password == "" {
		return dtos.SignUpOutput{}, shared.ErrInvalidSignUpInput
	}
	
	email := strings.ToLower(strings.TrimSpace(input.Email))
	parsedEmail, err := mail.ParseAddress(email)
	if err != nil || parsedEmail.Address != email {
		return dtos.SignUpOutput{}, shared.ErrInvalidSignUpInput
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return dtos.SignUpOutput{}, err
	}

	user := entities.NewUser(email, string(passwordHash))
	if err := i.userRepository.Create(ctx, user); err != nil {
		return dtos.SignUpOutput{}, err
	}

	return dtos.SignUpOutput{
		Id:        user.Id,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}, nil
}
