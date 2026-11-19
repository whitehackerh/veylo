package handlers

import (
	"io"
	"mime/multipart"
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/whitehackerh/veylo/api_server/src/generated/openapi"
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
	request := new(openapi.PostChannelRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, h.responsePresenter.ToErrorResponse(c.Request().URL.Path, "INVALID_REQUEST", "invalid request body"))
	}

	readFile := func(file *multipart.FileHeader) ([]byte, error) {
		if file == nil {
			return nil, nil
		}

		src, err := file.Open()
		if err != nil {
			return nil, err
		}
		defer src.Close()

		return io.ReadAll(src)
	}

	icon, err := readFile(request.Icon)
	if err != nil {
		return c.JSON(http.StatusBadRequest, h.responsePresenter.ToErrorResponse(c.Request().URL.Path, "INVALID_REQUEST", "invalid icon"))
	}

	banner, err := readFile(request.Banner)
	if err != nil {
		return c.JSON(http.StatusBadRequest, h.responsePresenter.ToErrorResponse(c.Request().URL.Path, "INVALID_REQUEST", "invalid banner"))
	}

	output, err := h.createChannelUseCase.Execute(c.Request().Context(), dtos.CreateChannelInput{
		UserId:      c.Get("user").(dtos.GetUserOutput).Id,
		Name:        request.Name,
		Handle:      request.Handle,
		Description: request.Description,
		Icon:        icon,
		Banner:      banner,
	})
	if err != nil {
		c.Logger().Error(err.Error())
		return c.JSON(http.StatusInternalServerError, h.responsePresenter.ToErrorResponse(c.Request().URL.Path, "INTERNAL_SERVER_ERROR", "failed"))
	}

	return c.JSON(http.StatusCreated, h.responsePresenter.ToApiResponse(c.Request().URL.Path, h.createChannelPresenter.ToResponse(&output), nil))
}
