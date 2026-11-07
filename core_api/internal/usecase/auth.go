package usecase

import (
	"errors"
	"strings"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

	"github.com/whitehackerh/veylo/api_server/internal/domain"
)

type TokenManager interface {
	Issue(userID string) (string, error)
	Verify(token string) (string, error)
}

type AuthUsecase struct {
	userRepository domain.UserRepository
	tokenManager   TokenManager
}

type AuthResult struct {
	User        *domain.User
	AccessToken string
}

func NewAuthUsecase(userRepository domain.UserRepository, tokenManager TokenManager) *AuthUsecase {
	return &AuthUsecase{userRepository: userRepository, tokenManager: tokenManager}
}

func (u *AuthUsecase) Register(email string, password string) (*domain.User, error) {
	email = strings.ToLower(strings.TrimSpace(email))
	if email == "" || len(password) < 8 {
		return nil, errors.New("email is required and password must be at least 8 characters")
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &domain.User{ID: uuid.NewString(), Email: email, PasswordHash: string(passwordHash)}
	if err := u.userRepository.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (u *AuthUsecase) Login(email string, password string) (*AuthResult, error) {
	email = strings.ToLower(strings.TrimSpace(email))
	user, err := u.userRepository.FindByEmail(email)
	if err != nil {
		return nil, domain.ErrInvalidCredentials
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		return nil, domain.ErrInvalidCredentials
	}

	accessToken, err := u.tokenManager.Issue(user.ID)
	if err != nil {
		return nil, err
	}

	return &AuthResult{User: user, AccessToken: accessToken}, nil
}

func (u *AuthUsecase) GetUser(userID string) (*domain.User, error) {
	return u.userRepository.FindByID(userID)
}
