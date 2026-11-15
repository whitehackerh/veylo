package postgres

import (
	"context"

	"github.com/whitehackerh/veylo/api_server/src/domain/entities"
)

type ChannelRepository struct {
	db DBTX
}

func NewChannelRepository(db DBTX) *ChannelRepository {
	return &ChannelRepository{db: db}
}

func (r *ChannelRepository) Create(ctx context.Context, channel *entities.Channel) error {
	_, err := r.db.Exec(ctx,
		`INSERT INTO channels (id, user_id, name, handle, description, icon_path, banner_path, created_at, updated_at, deleted_at)
		 VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`,
		channel.Id, channel.UserId, channel.Name, channel.Handle, channel.Description, channel.IconPath, channel.BannerPath, channel.CreatedAt, channel.UpdatedAt, channel.DeletedAt,
	)
	if err != nil {
		return err
	}
	return nil
}
