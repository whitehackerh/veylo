package ports

import (
	"context"

	"github.com/whitehackerh/veylo/api_server/src/usecases/dtos"
)

type RegisterUser interface {
	Execute(ctx context.Context, input dtos.RegisterUserInput) (*dtos.RegisterUserOutput, error)
}
