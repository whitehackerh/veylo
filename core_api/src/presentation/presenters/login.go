package presenters

import (
	"time"

	"github.com/whitehackerh/veylo/api_server/src/usecases/dtos"
)

func ToLoginData(output *dtos.LoginOutput) map[string]interface{} {
	return map[string]interface{}{
		"user": map[string]interface{}{
			"id":         output.ID.String(),
			"email":      output.Email,
			"created_at": output.CreatedAt.UTC().Truncate(time.Second),
		},
		"access_token": output.AccessToken,
		"token_type":   "Bearer",
	}
}
