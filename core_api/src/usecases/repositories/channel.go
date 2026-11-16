package repositories

import (
	"context"

	"github.com/whitehackerh/veylo/api_server/src/domain/entities"
)

type ChannelRepository interface {
	Create(ctx context.Context, channel *entities.Channel) error
}
