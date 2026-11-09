package handlers

import (
	"net/http"

	"github.com/labstack/echo/v5"

	"github.com/whitehackerh/veylo/api_server/src/domain/entities"
	"github.com/whitehackerh/veylo/api_server/src/presentation/openapi"
	"github.com/whitehackerh/veylo/api_server/src/presentation/presenters"
)

type GetUserHandler struct{}

func NewGetUserHandler() *GetUserHandler {
	return &GetUserHandler{}
}

func (h *GetUserHandler) Handle(c *echo.Context) error {
	user, ok := c.Get("user").(*entities.User)
	if !ok {
		return c.JSON(http.StatusUnauthorized, openapi.ErrorResponse{Error: "unauthorized"})
	}

	return c.JSON(http.StatusOK, presenters.ToGetUserResponse(user))
}
