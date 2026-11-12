package presenters

import (
	"time"

	"github.com/whitehackerh/veylo/api_server/src/generated/openapi"
)

func ToApiResponse(path string, data any, apiError *openapi.ApiError) openapi.ApiResponse {
	return openapi.ApiResponse{
		TimeStamp: time.Now().UTC(),
		Path:      path,
		Data:      data,
		Error:     apiError,
	}
}

func ToErrorResponse(path string, code string, message string) openapi.ApiResponse {
	return ToApiResponse(path, nil, &openapi.ApiError{
		Code:    code,
		Message: message,
	})
}
