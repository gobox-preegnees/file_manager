package file_usecase

import (
	"context"

	grpcController "github.com/gobox-preegnees/file_manager/internal/controller/grpc"
	entity "github.com/gobox-preegnees/file_manager/internal/domain/entity"
	storageDTO "github.com/gobox-preegnees/file_manager/internal/adapters/storage"
)

type IFileStorage interface {
	CreateOne(ctx context.Context, createReqDTO storageDTO.CreateReqDTO) (err error)
	FindOneByPath(ctx context.Context, findOneReqDTO storageDTO.FindOneReqDTO) (FindOneRespDTO storageDTO.FindOneRespDTO, err error)
	FindAll(ctx context.Context, findAllReqDTO storageDTO.FindAllReqDTO) (FindAllRespDTO storageDTO.FindAllRespDTO, err error)
	UpdateOne(ctx context.Context, updateOneReqDTO storageDTO.UpdateOneReqDTO) (err error)
	UpdateAll(ctx context.Context, updateAllReqDTO storageDTO.UpdateAllReqDTO) (err error)
	DeleteOne(ctx context.Context, deleteOneReqDTO storageDTO.DeleteOneReqDTO) (err error)
	DeleteAll(ctx context.Context, deleteAllReqDTO storageDTO.DeleteAllReqDTO) (err error)
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
		err := f.fileStorage.DeleteOne(ctx, storageDTO.DeleteOneReqDTO{
			Identifier: identifier,
			Path:       path,
		})
		if err != nil {
			return err
		}
	} else {
		err := f.fileStorage.DeleteAll(ctx, storageDTO.DeleteAllReqDTO{
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
	
	files, err := f.fileStorage.FindAll(ctx, storageDTO.FindAllReqDTO{
		Identifier: identifier,
	})
	if err != nil {
		return nil, err
	}
	return files.Files, nil
}

func (f *fileUsecase) RenameFile(ctx context.Context, identifier entity.Identifier, oldPath string, newPath string, hash string) error {
	
	if hash == "" {
		err := f.fileStorage.UpdateOne(ctx, storageDTO.UpdateOneReqDTO{
			Identifier: identifier,
			OldPath:    oldPath,
			NewPath:    newPath,
		})
		if err != nil {
			return err
		}
	} else {
		err := f.fileStorage.UpdateAll(ctx, storageDTO.UpdateAllReqDTO{
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
	
	err := f.fileStorage.CreateOne(ctx, storageDTO.CreateReqDTO{
		Identifier: identifier,
		File:       file,
	})
	if err != nil {
		return err
	}
	return nil
}

var _ grpcController.IFileUsecase = (*fileUsecase)(nil)
