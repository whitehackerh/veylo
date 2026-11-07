package handlers

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v5"

	"github.com/whitehackerh/veylo/api_server/src/presentation/openapi"
	"github.com/whitehackerh/veylo/api_server/src/presentation/presenters"
	"github.com/whitehackerh/veylo/api_server/src/usecases/dtos"
	"github.com/whitehackerh/veylo/api_server/src/usecases/ports"
	"github.com/whitehackerh/veylo/api_server/src/usecases/shared"
)

type RegisterUserHandler struct {
	registerUserUseCase ports.RegisterUserUseCase
}

func NewRegisterUserHandler(registerUserUseCase ports.RegisterUserUseCase) *RegisterUserHandler {
	return &RegisterUserHandler{registerUserUseCase: registerUserUseCase}
}

func (h *RegisterUserHandler) Handle(c *echo.Context) error {
	request := new(openapi.RegisterUserRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, openapi.ErrorResponse{Error: "invalid request body"})
	}

	output, err := h.registerUserUseCase.Execute(c.Request().Context(), dtos.RegisterUserInput{
		Email: request.Email,
		Password: request.Password,
	})
	if err != nil {
		if errors.Is(err, shared.ErrEmailAlreadyExists) {
			return c.JSON(http.StatusConflict, openapi.ErrorResponse{Error: "email already exists"})
		}
		if errors.Is(err, shared.ErrInvalidRegistrationInput) {
			return c.JSON(http.StatusBadRequest, openapi.ErrorResponse{Error: err.Error()})
		}
		return c.JSON(http.StatusInternalServerError, openapi.ErrorResponse{Error: "failed to register user"})
	}

	return c.JSON(http.StatusCreated, presenters.ToRegisterUserResponse(output))
}
