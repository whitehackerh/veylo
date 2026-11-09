package presenters

import (
	"time"

	"github.com/whitehackerh/veylo/api_server/src/presentation/openapi"
)

func ToCommonResponse(path string) openapi.CommonResponse {
	return openapi.CommonResponse{
		TimeStamp: time.Now().UTC(),
		Path:      path,
		Errors:    []openapi.ApiError{},
	}
}

func ToErrorResponse(path string, code string, message string) openapi.ErrorResponse {
	return openapi.ErrorResponse{
		Data: nil,
		CommonResponse: openapi.CommonResponse{
			TimeStamp: time.Now().UTC(),
			Path:      path,
			Errors: []openapi.ApiError{
				{
					Code:    code,
					Message: message,
				},
			},
		},
	}
}
