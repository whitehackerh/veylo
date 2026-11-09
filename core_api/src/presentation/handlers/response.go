package handlers

import (
	"time"

	"github.com/labstack/echo/v5"

	"github.com/whitehackerh/veylo/api_server/src/presentation/openapi"
)

func newApiResponse(c *echo.Context, data interface{}, apiError *openapi.ApiError) openapi.ApiResponse {
	return openapi.ApiResponse{
		TimeStamp: time.Now().UTC(),
		Path:      c.Request().URL.Path,
		Data:      data,
		Error:     apiError,
	}
}

func newErrorResponse(c *echo.Context, code string, message string) openapi.ApiResponse {
	return newApiResponse(c, nil, &openapi.ApiError{
		Code:    code,
		Message: message,
	})
}
