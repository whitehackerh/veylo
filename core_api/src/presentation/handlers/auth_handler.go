package handlers

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v5"

	"github.com/whitehackerh/veylo/api_server/src/presentation/presenters"
	"github.com/whitehackerh/veylo/api_server/src/usecases/dtos"
	"github.com/whitehackerh/veylo/api_server/src/usecases/ports"
	"github.com/whitehackerh/veylo/api_server/src/usecases/shared"
)

type AuthHandler struct {
	registerUser ports.RegisterUser
	login        ports.Login
	getUser      ports.GetUser
}

type credentialsRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func NewAuthHandler(registerUser ports.RegisterUser, login ports.Login, getUser ports.GetUser) *AuthHandler {
	return &AuthHandler{registerUser: registerUser, login: login, getUser: getUser}
}

func (h *AuthHandler) Register(c *echo.Context) error {
	request := new(credentialsRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request body"})
	}

	output, err := h.registerUser.Execute(c.Request().Context(), dtos.RegisterUserInput{
		Email: request.Email, Password: request.Password,
	})
	if err != nil {
		if errors.Is(err, shared.ErrEmailAlreadyExists) {
			return c.JSON(http.StatusConflict, map[string]string{"error": "email already exists"})
		}
		if errors.Is(err, shared.ErrInvalidRegistrationInput) {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to register user"})
	}
	return c.JSON(http.StatusCreated, presenters.ToUserResponse(output.User))
}

func (h *AuthHandler) Login(c *echo.Context) error {
	request := new(credentialsRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request body"})
	}

	output, err := h.login.Execute(c.Request().Context(), dtos.LoginInput{
		Email: request.Email, Password: request.Password,
	})
	if err != nil {
		if errors.Is(err, shared.ErrInvalidCredentials) {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "invalid email or password"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to login"})
	}
	return c.JSON(http.StatusOK, presenters.LoginResponse{
		User: presenters.ToUserResponse(output.User), AccessToken: output.AccessToken, TokenType: "Bearer",
	})
}

func (h *AuthHandler) Me(c *echo.Context) error {
	userID, ok := c.Get("user_id").(string)
	if !ok {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "unauthorized"})
	}

	output, err := h.getUser.Execute(c.Request().Context(), dtos.GetUserInput{UserID: userID})
	if err != nil {
		if errors.Is(err, shared.ErrUserNotFound) {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "unauthorized"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to get user"})
	}
	return c.JSON(http.StatusOK, presenters.ToUserResponse(output.User))
}
