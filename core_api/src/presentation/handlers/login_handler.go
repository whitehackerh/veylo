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

type LoginHandler struct {
	loginUseCase ports.LoginUseCase
}

type loginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func NewLoginHandler(loginUseCase ports.LoginUseCase) *LoginHandler {
	return &LoginHandler{loginUseCase: loginUseCase}
}

func (h *LoginHandler) Handle(c *echo.Context) error {
	request := new(loginRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request body"})
	}

	output, err := h.loginUseCase.Execute(c.Request().Context(), dtos.LoginInput{
		Email: request.Email, Password: request.Password,
	})
	if err != nil {
		if errors.Is(err, shared.ErrInvalidCredentials) {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "invalid email or password"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to login"})
	}

	return c.JSON(http.StatusOK, presenters.ToLoginResponse(output))
}
