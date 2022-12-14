package services

import (
	// "context"

	// entity "github.com/gobox-preegnees/file_manager/v1/internal/domain/entity"
)

type IFolderStorage interface {
	// Create(ctx context.Context, folder entity.Folder) (id int64, err error)
	// GetByPath(ctx context.Context, path entity.Path) (file entity.Folder, err error)
	// GetByModTime(ctx context.Context, modTime entity.ModTime) (file entity.Folder, err error)
	// UpdatePath(ctx context.Context, folder entity.Folder) (id int64, err error)
	// DeleteByPath(ctx context.Context, path entity.Path) (id int64, err error)
	// DeleteByModTime(ctx context.Context, modTime entity.ModTime) (id int64, err error)
}

type folderService struct {
	storage IFolderStorage
}

func NewFolderService(storage IFolderStorage) folderService {
	return folderService{
		storage: storage,
	}
}
