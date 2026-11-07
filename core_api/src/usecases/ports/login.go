package ports

import (
	"context"

	"github.com/whitehackerh/veylo/api_server/src/usecases/dtos"
)

type LoginUseCase interface {
	Execute(ctx context.Context, input dtos.LoginInput) (*dtos.LoginOutput, error)
}
