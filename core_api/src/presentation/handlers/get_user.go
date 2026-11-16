package handlers

import (
	"net/http"

	"github.com/labstack/echo/v5"

	"github.com/whitehackerh/veylo/api_server/src/presentation/presenters"
	"github.com/whitehackerh/veylo/api_server/src/usecases/dtos"
)

type GetUserHandler struct {
	responsePresenter *presenters.ResponsePresenter
	getUserPresenter  *presenters.GetUserPresenter
}

func NewGetUserHandler(
	responsePresenter *presenters.ResponsePresenter,
	getUserPresenter *presenters.GetUserPresenter,
) *GetUserHandler {
	return &GetUserHandler{
		responsePresenter: responsePresenter,
		getUserPresenter:  getUserPresenter,
	}
}

func (h *GetUserHandler) Handle(c *echo.Context) error {
	output, ok := c.Get("user").(dtos.GetUserOutput)
	if !ok {
		return c.JSON(http.StatusUnauthorized, h.responsePresenter.ToErrorResponse(c.Request().URL.Path, "UNAUTHORIZED", "unauthorized"))
	}

	return c.JSON(http.StatusOK, h.responsePresenter.ToApiResponse(c.Request().URL.Path, h.getUserPresenter.ToResponse(&output), nil))
}
