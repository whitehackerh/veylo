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

type RegisterUserHandler struct {
	registerUserUseCase ports.UseCase[dtos.RegisterUserInput, dtos.RegisterUserOutput]
}

func NewRegisterUserHandler(registerUserUseCase ports.UseCase[dtos.RegisterUserInput, dtos.RegisterUserOutput]) *RegisterUserHandler {
	return &RegisterUserHandler{registerUserUseCase: registerUserUseCase}
}

func (h *RegisterUserHandler) Handle(c *echo.Context) error {
	request := new(openapi.RegisterUserRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, presenters.ToErrorResponse(c.Request().URL.Path, "INVALID_REQUEST", "invalid request body"))
	}

	_, err := h.registerUserUseCase.Execute(c.Request().Context(), dtos.RegisterUserInput{
		Email:    request.Email,
		Password: request.Password,
	})
	if err != nil {
		if errors.Is(err, shared.ErrEmailAlreadyExists) {
			return c.JSON(http.StatusConflict, presenters.ToErrorResponse(c.Request().URL.Path, "EMAIL_ALREADY_EXISTS", "email already exists"))
		}
		if errors.Is(err, shared.ErrInvalidRegistrationInput) {
			return c.JSON(http.StatusBadRequest, presenters.ToErrorResponse(c.Request().URL.Path, "INVALID_REGISTRATION_INPUT", err.Error()))
		}
		return c.JSON(http.StatusInternalServerError, presenters.ToErrorResponse(c.Request().URL.Path, "INTERNAL_SERVER_ERROR", "failed to register user"))
	}

	return c.JSON(http.StatusCreated, presenters.ToAPIResponse(c.Request().URL.Path, presenters.ToRegisterUserResponse(), nil))
}
