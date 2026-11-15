package storages

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"syscall"

	"github.com/google/uuid"
	"github.com/whitehackerh/veylo/api_server/src/usecases/storages"
)

type LocalImageStorage struct {
	storagePath string
}

func NewLocalImageStorage(storagePath string) *LocalImageStorage {
	return &LocalImageStorage{
		storagePath: storagePath,
	}
}

func (s *LocalImageStorage) Save(ctx context.Context, target storages.ImageTarget, imageType storages.ImageType, id uuid.UUID, content []byte) (string, error) {
	if err := ctx.Err(); err != nil {
		return "", err
	}

	var relativePath string

	switch target {
	case storages.ImageTargetChannel:
		switch imageType {
		case storages.ImageTypeIcon, storages.ImageTypeBanner:
			relativePath = filepath.Join("images", "channels", id.String(), string(imageType)+".webp")
		default:
			return "", fmt.Errorf("unsupported image type: %s", imageType)
		}
	case storages.ImageTargetVideo:
		if imageType != storages.ImageTypeThumbnail {
			return "", fmt.Errorf("unsupported image type: %s", imageType)
		}

		relativePath = filepath.Join("images", "videos", id.String(), "thumbnail.webp")
	case storages.ImageTargetStream:
		if imageType != storages.ImageTypeThumbnail {
			return "", fmt.Errorf("unsupported image type: %s", imageType)
		}

		relativePath = filepath.Join("images", "streams", id.String(), "thumbnail.webp")
	default:
		return "", fmt.Errorf("unsupported image target: %s", target)
	}

	fullPath := filepath.Join(s.storagePath, relativePath)
	dirPath := filepath.Dir(fullPath)

	if err := os.MkdirAll(dirPath, 0777); err != nil {
		return "", err
	}

	if err := os.Chmod(dirPath, 0777); err != nil {
		return "", err
	}

	if err := os.WriteFile(fullPath, content, 0777); err != nil {
		return "", err
	}

	if err := os.Chmod(fullPath, 0777); err != nil {
		return "", err
	}

	return filepath.ToSlash(relativePath), nil
}

func (s *LocalImageStorage) Delete(ctx context.Context, path string) error {
	if err := ctx.Err(); err != nil {
		return err
	}

	fullPath := filepath.Join(s.storagePath, filepath.FromSlash(path))

	if err := os.Remove(fullPath); err != nil {
		if os.IsNotExist(err) {
			return nil
		}

		return err
	}

	if err := os.Remove(filepath.Dir(fullPath)); err != nil {
		if os.IsNotExist(err) || errors.Is(err, syscall.ENOTEMPTY) {
			return nil
		}

		return err
	}

	return nil
}
