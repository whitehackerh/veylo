package ports

import (
	"context"

	"github.com/whitehackerh/veylo/api_server/src/usecases/dtos"
)

type RegisterUserUseCase interface {
	Execute(ctx context.Context, input dtos.RegisterUserInput) (*dtos.RegisterUserOutput, error)
}
