package http

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v5"

	"github.com/whitehackerh/veylo/api_server/internal/domain"
	"github.com/whitehackerh/veylo/api_server/internal/usecase"
)

type AuthHandler struct {
	authUsecase *usecase.AuthUsecase
}

type credentialsRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type userResponse struct {
	ID        string `json:"id"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
}

type authResponse struct {
	User        userResponse `json:"user"`
	AccessToken string       `json:"access_token"`
	TokenType   string       `json:"token_type"`
}

func NewAuthHandler(authUsecase *usecase.AuthUsecase) *AuthHandler {
	return &AuthHandler{authUsecase: authUsecase}
}

func (h *AuthHandler) Register(c *echo.Context) error {
	request := new(credentialsRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request body"})
	}
	user, err := h.authUsecase.Register(request.Email, request.Password)
	if err != nil {
		if errors.Is(err, domain.ErrEmailAlreadyExists) {
			return c.JSON(http.StatusConflict, map[string]string{"error": "email already exists"})
		}
		if err.Error() == "email is required and password must be at least 8 characters" {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to register user"})
	}
	return c.JSON(http.StatusCreated, toUserResponse(user))
}

func (h *AuthHandler) Login(c *echo.Context) error {
	request := new(credentialsRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request body"})
	}
	result, err := h.authUsecase.Login(request.Email, request.Password)
	if err != nil {
		if errors.Is(err, domain.ErrInvalidCredentials) {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "invalid email or password"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to login"})
	}
	return c.JSON(http.StatusOK, authResponse{
		User: toUserResponse(result.User), AccessToken: result.AccessToken, TokenType: "Bearer",
	})
}

func (h *AuthHandler) Me(c *echo.Context) error {
	userID, ok := c.Get("user_id").(string)
	if !ok {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "unauthorized"})
	}
	user, err := h.authUsecase.GetUser(userID)
	if err != nil {
		if errors.Is(err, domain.ErrUserNotFound) {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "unauthorized"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to get user"})
	}
	return c.JSON(http.StatusOK, toUserResponse(user))
}

func toUserResponse(user *domain.User) userResponse {
	return userResponse{
		ID: user.ID, Email: user.Email,
		CreatedAt: user.CreatedAt.UTC().Format("2006-01-02T15:04:05Z07:00"),
	}
}
