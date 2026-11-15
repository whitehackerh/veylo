package entities

import (
	"time"

	"github.com/google/uuid"
)

type Channel struct {
	Id          uuid.UUID
	UserId      uuid.UUID
	Name        string
	Handle      string
	Description *string
	IconPath    *string
	BannerPath  *string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
}

func NewChannel(
	id uuid.UUID,
	user_id uuid.UUID,
	name string,
	handle string,
	description *string,
	iconPath *string,
	bannerPath *string,
) *Channel {
	now := time.Now()
	return &Channel{
		Id:          id,
		UserId:      user_id,
		Name:        name,
		Handle:      handle,
		Description: description,
		IconPath:    iconPath,
		BannerPath:  bannerPath,
		CreatedAt:   now,
		UpdatedAt:   now,
		DeletedAt:   nil,
	}
}

func RestoreChannel(
	id uuid.UUID,
	name string,
	handle string,
	description *string,
	iconPath *string,
	bannerPath *string,
	createdAt time.Time,
	updatedAt time.Time,
	deletedAt *time.Time,
) *Channel {
	return &Channel{
		Id:          id,
		Name:        name,
		Handle:      handle,
		Description: description,
		IconPath:    iconPath,
		BannerPath:  bannerPath,
		CreatedAt:   createdAt,
		UpdatedAt:   updatedAt,
		DeletedAt:   deletedAt,
	}
}
