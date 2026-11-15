package storages

import (
	"context"

	"github.com/google/uuid"
)

type ImageTarget string

const (
	ImageTargetChannel ImageTarget = "channel"
	ImageTargetVideo   ImageTarget = "video"
	ImageTargetStream  ImageTarget = "stream"
)

type ImageType string

const (
	ImageTypeIcon      ImageType = "icon"
	ImageTypeBanner    ImageType = "banner"
	ImageTypeThumbnail ImageType = "thumbnail"
)

type ImageStorage interface {
	Save(ctx context.Context, target ImageTarget, imageType ImageType, id uuid.UUID, content []byte) (string, error)
	Delete(ctx context.Context, path string) error
}
