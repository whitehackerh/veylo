package presenters

import (
	"time"

	"github.com/whitehackerh/veylo/api_server/src/usecases/dtos"
)

func ToRegisterUserData(output *dtos.RegisterUserOutput) map[string]interface{} {
	return map[string]interface{}{
		"id":         output.ID.String(),
		"email":      output.Email,
		"created_at": output.CreatedAt.UTC().Truncate(time.Second),
	}
}
