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

type SignInHandler struct {
	responsePresenter *presenters.ResponsePresenter
	signInUseCase     ports.UseCase[dtos.SignInInput, dtos.SignInOutput]
	signInPresenter   *presenters.SignInPresenter
}

func NewSignInHandler(
	responsePresenter *presenters.ResponsePresenter,
	signInUseCase ports.UseCase[dtos.SignInInput, dtos.SignInOutput],
	signInPresenter *presenters.SignInPresenter,
) *SignInHandler {
	return &SignInHandler{
		responsePresenter: responsePresenter,
		signInUseCase:     signInUseCase,
		signInPresenter:   signInPresenter,
	}
}

func (h *SignInHandler) Handle(c *echo.Context) error {
	request := new(openapi.SignInRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, h.responsePresenter.ToErrorResponse(c.Request().URL.Path, "INVALID_REQUEST", "invalid request body"))
	}

	output, err := h.signInUseCase.Execute(c.Request().Context(), dtos.SignInInput{
		Email:    string(request.Email),
		Password: request.Password,
	})
	if err != nil {
		if errors.Is(err, shared.ErrInvalidCredentials) {
			return c.JSON(http.StatusUnauthorized, h.responsePresenter.ToErrorResponse(c.Request().URL.Path, "INVALID_CREDENTIALS", "invalid email or password"))
		}
		return c.JSON(http.StatusInternalServerError, h.responsePresenter.ToErrorResponse(c.Request().URL.Path, "INTERNAL_SERVER_ERROR", "failed to sign in"))
	}

	c.Response().Header().Set("Authorization", "Bearer "+output.AccessToken)
	return c.JSON(http.StatusOK, h.responsePresenter.ToApiResponse(c.Request().URL.Path, h.signInPresenter.ToResponse(), nil))
}
