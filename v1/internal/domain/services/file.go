package services

import (
	"context"

	entity "github.com/gobox-preegnees/file_manager/v1/internal/domain/entity"
)

type IFileStorage interface {
	GetByPath(ctx context.Context, path entity.Path) (file entity.File, err error)
	Create(ctx context.Context, file entity.File) (err error)
	Update(ctx context.Context, file entity.File) (id int64, err error)
	DeleteByPath(ctx context.Context, path entity.Path) (id int64, err error)
}

type fileService struct {
	storage IFileStorage
}

func NewFileService(storage IFileStorage) fileService {
	return fileService{
		storage: storage,
	}
}

func (f fileService) SaveFile(ctx context.Context, file entity.File) error {
	return f.storage.Create(ctx, file)
}
