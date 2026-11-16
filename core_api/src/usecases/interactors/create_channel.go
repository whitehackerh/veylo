package interactors

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/whitehackerh/veylo/api_server/src/domain/entities"
	"github.com/whitehackerh/veylo/api_server/src/usecases/dtos"
	"github.com/whitehackerh/veylo/api_server/src/usecases/repositories"
	"github.com/whitehackerh/veylo/api_server/src/usecases/storages"
)

type CreateChannelInteractor struct {
	imageStorage storages.ImageStorage
	uow          repositories.UnitOfWork
}

func NewCreateChannelInteractor(
	imageStorage storages.ImageStorage,
	uow repositories.UnitOfWork,
) *CreateChannelInteractor {
	return &CreateChannelInteractor{
		imageStorage: imageStorage,
		uow:          uow,
	}
}

func (i *CreateChannelInteractor) Execute(
	ctx context.Context,
	input dtos.CreateChannelInput,
) (dtos.CreateChannelOutput, error) {
	channelId := uuid.New()

	type imageSaveInput struct {
		imageType storages.ImageType
		content   []byte
	}
	images := []imageSaveInput{
		{imageType: storages.ImageTypeIcon, content: input.Icon},
		{imageType: storages.ImageTypeBanner, content: input.Banner},
	}
	imagePaths := make(map[storages.ImageType]string, len(images))

	var wg sync.WaitGroup
	var mu sync.Mutex
	errCh := make(chan error, len(images))

	for _, image := range images {
		if image.content == nil {
			continue
		}

		wg.Add(1)

		go func(image imageSaveInput) {
			defer wg.Done()

			path, err := i.imageStorage.Save(ctx, storages.ImageTargetChannel, image.imageType, channelId, image.content)
			if err != nil {
				errCh <- err
				return
			}

			mu.Lock()
			imagePaths[image.imageType] = path
			mu.Unlock()
		}(image)
	}

	wg.Wait()
	close(errCh)

	for err := range errCh {
		if err != nil {
			return dtos.CreateChannelOutput{}, errors.Join(err, i.cleanupImages(imagePaths))
		}
	}

	var iconPath *string
	if path, ok := imagePaths[storages.ImageTypeIcon]; ok {
		iconPath = &path
	}

	var bannerPath *string
	if path, ok := imagePaths[storages.ImageTypeBanner]; ok {
		bannerPath = &path
	}

	channel := entities.NewChannel(
		channelId,
		input.UserId,
		input.Name,
		input.Handle,
		input.Description,
		iconPath,
		bannerPath,
	)

	err := i.uow.Execute(ctx, func(repos repositories.Repositories) error {
		return repos.Channel().Create(ctx, channel)
	})
	if err != nil {
		return dtos.CreateChannelOutput{}, errors.Join(err, i.cleanupImages(imagePaths))
	}

	return dtos.CreateChannelOutput{
		Id:          channel.Id,
		Name:        channel.Name,
		Handle:      channel.Handle,
		Description: channel.Description,
		IconPath:    channel.IconPath,
		BannerPath:  channel.BannerPath,
		CreatedAt:   channel.CreatedAt,
		UpdatedAt:   channel.UpdatedAt,
	}, nil
}

func (i *CreateChannelInteractor) cleanupImages(imagePaths map[storages.ImageType]string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var cleanupErrors []error

	for _, path := range imagePaths {
		if err := i.imageStorage.Delete(ctx, path); err != nil {
			cleanupErrors = append(cleanupErrors, err)
		}
	}

	return errors.Join(cleanupErrors...)
}
