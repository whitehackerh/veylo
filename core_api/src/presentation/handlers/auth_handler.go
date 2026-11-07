package handlers

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v5"

	"github.com/whitehackerh/veylo/api_server/src/presentation/presenters"
	"github.com/whitehackerh/veylo/api_server/src/usecases/dtos"
	"github.com/whitehackerh/veylo/api_server/src/usecases/interactors"
	"github.com/whitehackerh/veylo/api_server/src/usecases/shared"
)

type AuthHandler struct {
	registerUserInteractor *interactors.RegisterUserInteractor
	loginInteractor        *interactors.LoginInteractor
	getUserInteractor      *interactors.GetUserInteractor
}

type credentialsRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func NewAuthHandler(
	registerUserInteractor *interactors.RegisterUserInteractor,
	loginInteractor *interactors.LoginInteractor,
	getUserInteractor *interactors.GetUserInteractor,
) *AuthHandler {
	return &AuthHandler{
		registerUserInteractor: registerUserInteractor,
		loginInteractor:        loginInteractor,
		getUserInteractor:      getUserInteractor,
	}
}

func (h *AuthHandler) Register(c *echo.Context) error {
	request := new(credentialsRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request body"})
	}

	user, err := h.registerUserInteractor.Execute(c.Request().Context(), request.Email, request.Password)
	if err != nil {
		if errors.Is(err, shared.ErrEmailAlreadyExists) {
			return c.JSON(http.StatusConflict, map[string]string{"error": "email already exists"})
		}
		if errors.Is(err, shared.ErrInvalidRegistrationInput) {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to register user"})
	}

	return c.JSON(http.StatusCreated, presenters.ToUserResponse(user))
}

func (h *AuthHandler) Login(c *echo.Context) error {
	request := new(credentialsRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request body"})
	}

	result, err := h.loginInteractor.Execute(c.Request().Context(), dtos.LoginInput{
		Email:    request.Email,
		Password: request.Password,
	})
	if err != nil {
		if errors.Is(err, shared.ErrInvalidCredentials) {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "invalid email or password"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to login"})
	}

	return c.JSON(http.StatusOK, presenters.LoginResponse{
		User:        presenters.ToUserResponse(result.User),
		AccessToken: result.AccessToken,
		TokenType:   "Bearer",
	})
}

func (h *AuthHandler) Me(c *echo.Context) error {
	userID, ok := c.Get("user_id").(string)
	if !ok {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "unauthorized"})
	}

	user, err := h.getUserInteractor.Execute(c.Request().Context(), userID)
	if err != nil {
		if errors.Is(err, shared.ErrUserNotFound) {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "unauthorized"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to get user"})
	}

	return c.JSON(http.StatusOK, presenters.ToUserResponse(user))
}
