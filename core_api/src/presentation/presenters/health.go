package presenters

import "github.com/whitehackerh/veylo/api_server/src/presentation/openapi"

func ToHealthData() openapi.HealthResponse {
	return openapi.HealthResponse{Status: "ok"}
}
