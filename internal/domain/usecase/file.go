package file_usecase

import (
	"context"

	grpc_controller "github.com/gobox-preegnees/file_manager/internal/controller/grpc"
	entity "github.com/gobox-preegnees/file_manager/internal/domain/entity"
)

type UpdateOneReqDTO struct {
	entity.Identifier
	OldPath string
	NewPath string
}
type UpdateMoreReqDTO struct {
	entity.Identifier
	OldPath string
	NewPath string
}
type CreateReqOne struct {
	entity.Identifier
	entity.File
}
type ReadMoreReqDTO struct {
	entity.Identifier
}
type ReadMoreRespDTO struct {
	Files []entity.File
}
type DeleteOneDTO struct {
	entity.Identifier
	Path string
}
type DeleteMoreDTO struct {
	entity.Identifier
	Path string
}

type IFileStorage interface {
	CreateOne(ctx context.Context, createReqOne CreateReqOne) (err error)
	ReadMore(ctx context.Context, readMoreReqDTO ReadMoreReqDTO) (readMoreRespDTO ReadMoreRespDTO, err error)
	UpdateOne(ctx context.Context, updateOneReqDTO UpdateOneReqDTO) (err error)
	UpdateMore(ctx context.Context, updateMoreReqDTO UpdateMoreReqDTO) (err error)
	DeleteOne(ctx context.Context, deleteOne DeleteOneDTO) (err error)
	DeleteMore(ctx context.Context, deleteMore DeleteMoreDTO) (err error)
}

type fileUsecase struct {
	fileStorage IFileStorage
}

func New(fileStorage IFileStorage) fileUsecase {
	return fileUsecase{
		fileStorage: fileStorage,
	}
}

func (f *fileUsecase) DeleteFile(ctx context.Context, identifier entity.Identifier, path string, hash string) error {
	if hash == "" {
		err := f.fileStorage.DeleteMore(ctx, DeleteMoreDTO{
			Identifier: identifier,
			Path:       path,
		})
		if err != nil {
			return err
		}
	} else {
		err := f.fileStorage.DeleteMore(ctx, DeleteMoreDTO{
			Identifier: identifier,
			Path:       path,
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func (f *fileUsecase) GetFiles(ctx context.Context, identifier entity.Identifier) ([]entity.File, error) {
	files, err := f.fileStorage.ReadMore(ctx, ReadMoreReqDTO{
		Identifier: identifier,
	})
	if err != nil {
		return nil, err
	}
	return files.Files, nil
}

func (f *fileUsecase) RenameFile(ctx context.Context, identifier entity.Identifier, oldPath string, newPath string, hash string) error {
	if hash == "" {
		err := f.fileStorage.UpdateOne(ctx, UpdateOneReqDTO{
			Identifier: identifier,
			OldPath:    oldPath,
			NewPath:    newPath,
		})
		if err != nil {
			return err
		}
	} else {
		err := f.fileStorage.UpdateMore(ctx, UpdateMoreReqDTO{
			Identifier: identifier,
			OldPath:    oldPath,
			NewPath:    newPath,
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func (f *fileUsecase) SaveFile(ctx context.Context, identifier entity.Identifier, file entity.File) error {
	err := f.fileStorage.CreateOne(ctx, CreateReqOne{
		Identifier: identifier,
		File:       file,
	})
	if err != nil {
		return err
	}
	return nil
}

var _ grpc_controller.IFileUsecase = (*fileUsecase)(nil)
