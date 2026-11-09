package presenters

import (
	"time"

	"github.com/whitehackerh/veylo/api_server/src/usecases/dtos"
)

func ToGetUserData(output *dtos.GetUserOutput) map[string]interface{} {
	return map[string]interface{}{
		"id":         output.ID.String(),
		"email":      output.Email,
		"created_at": output.CreatedAt.UTC().Truncate(time.Second),
	}
}
