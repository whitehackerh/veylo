package handlers

import (
	"net/http"

	"github.com/labstack/echo/v5"

	"github.com/whitehackerh/veylo/api_server/src/presentation/presenters"
	"github.com/whitehackerh/veylo/api_server/src/usecases/dtos"
)

type GetUserHandler struct{}

func NewGetUserHandler() *GetUserHandler {
	return &GetUserHandler{}
}

func (h *GetUserHandler) Handle(c *echo.Context) error {
	output, ok := c.Get("user").(*dtos.GetUserOutput)
	if !ok {
		return c.JSON(http.StatusUnauthorized, presenters.ToErrorResponse(c.Request().URL.Path, "UNAUTHORIZED", "unauthorized"))
	}

	return c.JSON(http.StatusOK, presenters.ToResponse(presenters.ToGetUserData(output), c.Request().URL.Path))
}
