package file_usecase

import (
	"context"

	grpcController "github.com/gobox-preegnees/file_manager/internal/controller/grpc"
	entity "github.com/gobox-preegnees/file_manager/internal/domain/entity"
	repoDTO "github.com/gobox-preegnees/file_manager/internal/adapters/repo"
)

type FileRepo interface {
	FindOneByPath(ctx context.Context, username, folderID, path string) (repoDTO.FileDTO, error)
	FindAll(ctx context.Context, username, folderID string) ([]repoDTO.FileDTO, error)
	Save(ctx context.Context, file repoDTO.FileDTO) (error)
}

type fileUsecase struct {
	fileRepo FileRepo
}

func New(fileRepo FileRepo) fileUsecase {
	
	return fileUsecase{
		fileRepo: fileRepo,
	}
}

func (f *fileUsecase) SaveFile(ctx context.Context, identifier entity.Identifier, file entity.File) error {
	
	return nil
}

func (f *fileUsecase) DeleteFile(ctx context.Context, identifier entity.Identifier, path string, hash string) error {
	
	if hash == "" {
		
	} else {
		
	}
	return nil
}

func (f *fileUsecase) GetFiles(ctx context.Context, identifier entity.Identifier) ([]entity.File, error) {
	
	
	return []entity.File{}, nil
}

func (f *fileUsecase) RenameFile(ctx context.Context, identifier entity.Identifier, oldPath string, newPath string, hash string) error {
	
	if hash == "" {
		
	} else {
		
	}
	return nil
}

var _ grpcController.IFileUsecase = (*fileUsecase)(nil)
