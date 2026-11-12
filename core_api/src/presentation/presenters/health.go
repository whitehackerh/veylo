package presenters

import "github.com/whitehackerh/veylo/api_server/src/generated/openapi"

func ToHealthData() openapi.HealthResponse {
	return openapi.HealthResponse{Status: "ok"}
}
