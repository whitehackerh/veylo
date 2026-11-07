package ports

import (
	"context"

	"github.com/whitehackerh/veylo/api_server/src/usecases/dtos"
)

type GetUserUseCase interface {
	Execute(ctx context.Context, input dtos.GetUserInput) (*dtos.GetUserOutput, error)
}
