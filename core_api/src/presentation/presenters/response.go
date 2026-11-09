package presenters

import (
	"time"

	"github.com/whitehackerh/veylo/api_server/src/presentation/openapi"
)

func ToResponse(data map[string]interface{}, path string) openapi.ApiResponse {
	return openapi.ApiResponse{
		Data:      data,
		TimeStamp: time.Now().UTC(),
		Path:      path,
		Errors:    []openapi.ApiError{},
	}
}

func ToErrorResponse(path string, code string, message string) openapi.ApiResponse {
	return openapi.ApiResponse{
		Data:      nil,
		TimeStamp: time.Now().UTC(),
		Path:      path,
		Errors: []openapi.ApiError{
			{
				Code:    code,
				Message: message,
			},
		},
	}
}
