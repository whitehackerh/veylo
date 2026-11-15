package dtos

import (
	"time"

	"github.com/google/uuid"
)

type CreateChannelInput struct {
	UserId      uuid.UUID
	Name        string
	Handle      string
	Description *string
	Icon        []byte
	Banner      []byte
}

type CreateChannelOutput struct {
	Id          uuid.UUID
	Name        string
	Handle      string
	Description *string
	IconPath    *string
	BannerPath  *string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
