package handlers

import (
	"errors"
	"io"
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/whitehackerh/veylo/api_server/src/presentation/presenters"
	"github.com/whitehackerh/veylo/api_server/src/usecases/dtos"
	"github.com/whitehackerh/veylo/api_server/src/usecases/ports"
)

type CreateChannelHandler struct {
	responsePresenter      *presenters.ResponsePresenter
	createChannelUseCase   ports.UseCase[dtos.CreateChannelInput, dtos.CreateChannelOutput]
	createChannelPresenter *presenters.CreateChannelPresenter
}

func NewCreateChannelHandler(
	responsePresenter *presenters.ResponsePresenter,
	createChannelUseCase ports.UseCase[dtos.CreateChannelInput, dtos.CreateChannelOutput],
	createChannelPresenter *presenters.CreateChannelPresenter,
) *CreateChannelHandler {
	return &CreateChannelHandler{
		responsePresenter:      responsePresenter,
		createChannelUseCase:   createChannelUseCase,
		createChannelPresenter: createChannelPresenter,
	}
}

func (h *CreateChannelHandler) Handle(c *echo.Context) error {
	if err := c.Request().ParseMultipartForm(32 << 20); err != nil {
		return c.JSON(http.StatusBadRequest, h.responsePresenter.ToErrorResponse(c.Request().URL.Path, "INVALID_REQUEST", "invalid request body"))
	}

	var description *string
	if values := c.Request().MultipartForm.Value["description"]; len(values) > 0 {
		description = &values[0]
	}

	readFile := func(field string) ([]byte, error) {
		file, err := c.FormFile(field)
		if err != nil {
			if errors.Is(err, http.ErrMissingFile) {
				return nil, nil
			}
			return nil, err
		}

		src, err := file.Open()
		if err != nil {
			return nil, err
		}
		defer src.Close()

		return io.ReadAll(src)
	}

	icon, err := readFile("icon")
	if err != nil {
		return c.JSON(http.StatusBadRequest, h.responsePresenter.ToErrorResponse(c.Request().URL.Path, "INVALID_REQUEST", "invalid icon"))
	}

	banner, err := readFile("banner")
	if err != nil {
		return c.JSON(http.StatusBadRequest, h.responsePresenter.ToErrorResponse(c.Request().URL.Path, "INVALID_REQUEST", "invalid banner"))
	}

	output, err := h.createChannelUseCase.Execute(c.Request().Context(), dtos.CreateChannelInput{
		UserId:      c.Get("user").(dtos.GetUserOutput).Id,
		Name:        c.FormValue("name"),
		Handle:      c.FormValue("handle"),
		Description: description,
		Icon:        icon,
		Banner:      banner,
	})
	if err != nil {
		c.Logger().Error(err.Error())
		return c.JSON(http.StatusInternalServerError, h.responsePresenter.ToErrorResponse(c.Request().URL.Path, "INTERNAL_SERVER_ERROR", "failed"))
	}

	return c.JSON(http.StatusCreated, h.responsePresenter.ToApiResponse(c.Request().URL.Path, h.createChannelPresenter.ToResponse(&output), nil))
}
