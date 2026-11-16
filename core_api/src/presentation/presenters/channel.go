package presenters

import (
	"time"

	"github.com/google/uuid"
	openapi_types "github.com/oapi-codegen/runtime/types"

	"github.com/whitehackerh/veylo/api_server/src/generated/openapi"
)

type Channel struct {
	Id          uuid.UUID
	Name        string
	Handle      string
	Description *string
	IconPath    *string
	BannerPath  *string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type ChannelPresenter struct {
	mediaBaseUrl string
}

func NewChannelPresenter(mediaBaseUrl string) *ChannelPresenter {
	return &ChannelPresenter{
		mediaBaseUrl: mediaBaseUrl,
	}
}

func (p *ChannelPresenter) ToResponse(channel Channel) openapi.Channel {
	return openapi.Channel{
		BannerUrl:   p.toMediaUrl(channel.BannerPath),
		CreatedAt:   channel.CreatedAt.UTC().Truncate(time.Second),
		Description: channel.Description,
		Handle:      channel.Handle,
		IconUrl:     p.toMediaUrl(channel.IconPath),
		Id:          openapi_types.UUID(channel.Id),
		Name:        channel.Name,
		UpdatedAt:   channel.UpdatedAt.UTC().Truncate(time.Second),
	}
}

func (p *ChannelPresenter) toMediaUrl(path *string) *string {
	if path == nil {
		return nil
	}

	mediaUrl := p.mediaBaseUrl + "/" + *path
	return &mediaUrl
}
