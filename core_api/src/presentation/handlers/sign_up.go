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

type SignUpHandler struct {
	responsePresenter *presenters.ResponsePresenter
	signUpUseCase     ports.UseCase[dtos.SignUpInput, dtos.SignUpOutput]
	signUpPresenter   *presenters.SignUpPresenter
}

func NewSignUpHandler(
	responsePresenter *presenters.ResponsePresenter,
	signUpUseCase ports.UseCase[dtos.SignUpInput, dtos.SignUpOutput],
	signUpPresenter *presenters.SignUpPresenter,
) *SignUpHandler {
	return &SignUpHandler{signUpUseCase: signUpUseCase}
}

func (h *SignUpHandler) Handle(c *echo.Context) error {
	request := new(openapi.SignUpRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, h.responsePresenter.ToErrorResponse(c.Request().URL.Path, "INVALID_REQUEST", "invalid request body"))
	}

	_, err := h.signUpUseCase.Execute(c.Request().Context(), dtos.SignUpInput{
		Email:    string(request.Email),
		Password: request.Password,
		Name:     request.Name,
	})
	if err != nil {
		if errors.Is(err, shared.ErrEmailAlreadyExists) {
			return c.JSON(http.StatusConflict, h.responsePresenter.ToErrorResponse(c.Request().URL.Path, "EMAIL_ALREADY_EXISTS", "email already exists"))
		}
		if errors.Is(err, shared.ErrInvalidSignUpInput) {
			return c.JSON(http.StatusBadRequest, h.responsePresenter.ToErrorResponse(c.Request().URL.Path, "INVALID_SIGN_UP_INPUT", err.Error()))
		}
		return c.JSON(http.StatusInternalServerError, h.responsePresenter.ToErrorResponse(c.Request().URL.Path, "INTERNAL_SERVER_ERROR", "failed to sign up"))
	}

	return c.JSON(http.StatusCreated, h.responsePresenter.ToApiResponse(c.Request().URL.Path, h.signUpPresenter.ToResponse(), nil))
}
