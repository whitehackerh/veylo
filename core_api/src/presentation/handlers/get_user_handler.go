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

type GetUserHandler struct {
	getUserUseCase ports.GetUserUseCase
}

func NewGetUserHandler(getUserUseCase ports.GetUserUseCase) *GetUserHandler {
	return &GetUserHandler{getUserUseCase: getUserUseCase}
}

func (h *GetUserHandler) Handle(c *echo.Context) error {
	userID, ok := c.Get("user_id").(string)
	if !ok {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "unauthorized"})
	}

	output, err := h.getUserUseCase.Execute(c.Request().Context(), dtos.GetUserInput{UserID: userID})
	if err != nil {
		if errors.Is(err, shared.ErrUserNotFound) {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "unauthorized"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to get user"})
	}

	return c.JSON(http.StatusOK, presenters.ToGetUserResponse(output))
}
