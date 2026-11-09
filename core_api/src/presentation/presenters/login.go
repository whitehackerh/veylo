package presenters

import (
	"github.com/whitehackerh/veylo/api_server/src/presentation/openapi"
)

func ToLoginResponse(path string) openapi.LoginResponse {
	return openapi.LoginResponse{
		Data:            struct{}{},
		CommonResponse: ToCommonResponse(path),
	}
}
