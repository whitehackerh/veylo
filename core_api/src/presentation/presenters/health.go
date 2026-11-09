package presenters

import "github.com/whitehackerh/veylo/api_server/src/presentation/openapi"

func ToHealthResponse() openapi.HealthResponse {
	return openapi.HealthResponse{Status: "ok"}
}
