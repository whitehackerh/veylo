package handlers

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v5"

	"github.com/whitehackerh/veylo/api_server/src/generated/openapi"
	"github.com/whitehackerh/veylo/api_server/src/presentation/presenters"
	"github.com/whitehackerh/veylo/api_server/src/usecases/dtos"
	"github.com/whitehackerh/veylo/api_server/src/usecases/ports"
	"github.com/whitehackerh/veylo/api_server/src/usecases/shared"
)

type LoginHandler struct {
	loginUseCase ports.UseCase[dtos.LoginInput, dtos.LoginOutput]
}

func NewLoginHandler(loginUseCase ports.UseCase[dtos.LoginInput, dtos.LoginOutput]) *LoginHandler {
	return &LoginHandler{loginUseCase: loginUseCase}
}

func (h *LoginHandler) Handle(c *echo.Context) error {
	request := new(openapi.LoginRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, presenters.ToErrorResponse(c.Request().URL.Path, "INVALID_REQUEST", "invalid request body"))
	}

	output, err := h.loginUseCase.Execute(c.Request().Context(), dtos.LoginInput{
		Email:    request.Email,
		Password: request.Password,
	})
	if err != nil {
		if errors.Is(err, shared.ErrInvalidCredentials) {
			return c.JSON(http.StatusUnauthorized, presenters.ToErrorResponse(c.Request().URL.Path, "INVALID_CREDENTIALS", "invalid email or password"))
		}
		return c.JSON(http.StatusInternalServerError, presenters.ToErrorResponse(c.Request().URL.Path, "INTERNAL_SERVER_ERROR", "failed to login"))
	}

	c.Response().Header().Set("Authorization", "Bearer "+output.AccessToken)
	return c.JSON(http.StatusOK, presenters.ToAPIResponse(c.Request().URL.Path, presenters.ToLoginResponse(), nil))
}
