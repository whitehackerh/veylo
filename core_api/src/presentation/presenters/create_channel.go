package presenters

import (
	"github.com/whitehackerh/veylo/api_server/src/generated/openapi"
	"github.com/whitehackerh/veylo/api_server/src/usecases/dtos"
)

type CreateChannelPresenter struct {
	channelPresenter *ChannelPresenter
}

func NewCreateChannelPresenter(channelPresenter *ChannelPresenter) *CreateChannelPresenter {
	return &CreateChannelPresenter{
		channelPresenter: channelPresenter,
	}
}

func (p *CreateChannelPresenter) ToResponse(output *dtos.CreateChannelOutput) openapi.Channel {
	return p.channelPresenter.ToResponse(Channel{
		Id:          output.Id,
		Name:        output.Name,
		Handle:      output.Handle,
		Description: output.Description,
		IconPath:    output.IconPath,
		BannerPath:  output.BannerPath,
		CreatedAt:   output.CreatedAt,
		UpdatedAt:   output.UpdatedAt,
	})
}
