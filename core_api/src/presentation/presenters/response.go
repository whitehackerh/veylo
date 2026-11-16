package presenters

import (
	"time"

	"github.com/whitehackerh/veylo/api_server/src/generated/openapi"
)

type ResponsePresenter struct{}

func NewResponsePresenter() *ResponsePresenter {
	return &ResponsePresenter{}
}

func (p *ResponsePresenter) ToApiResponse(path string, data any, apiError *openapi.ApiError) openapi.ApiResponse {
	return openapi.ApiResponse{
		TimeStamp: time.Now().UTC(),
		Path:      path,
		Data:      data,
		Error:     apiError,
	}
}

func (p *ResponsePresenter) ToErrorResponse(path string, code string, message string) openapi.ApiResponse {
	return p.ToApiResponse(path, nil, &openapi.ApiError{
		Code:    code,
		Message: message,
	})
}
