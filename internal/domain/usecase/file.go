package file_usecase

import (
	"context"

	repoDTO "github.com/gobox-preegnees/file_manager/internal/adapters/repo"
	grpcController "github.com/gobox-preegnees/file_manager/internal/controller/grpc"
	entity "github.com/gobox-preegnees/file_manager/internal/domain/entity"
)

type FileRepo interface {
	SaveOwner(ctx context.Context, saveOwnerDTO repoDTO.SaveOwnerDTO) (int, error)
	RenameOwner(ctx context.Context, renameOwnerDTO repoDTO.RenameOwnerDTO) error
	DeleteOwner(ctx context.Context, deleteOwnerDTO repoDTO.DeleteOwnerDTO) error
	FindAllOwners(ctx context.Context, findAllOwnersReqDTO repoDTO.FindAllOwnersReqDTO) (repoDTO.FindAllOwnersRespDTO, error)
	SaveFile(ctx context.Context, saveFileReqDTO repoDTO.SaveFileReqDTO) (int, error)
	SetState(ctx context.Context, setStateReqDTO repoDTO.SetStateReqDTO) error
	RenameFile(ctx context.Context, renameFileReqDTO repoDTO.RenameFileReqDTO) error
	DeleteFile(ctx context.Context, deleteFileReqDTO repoDTO.DeleteFileReqDTO) error
	RestoreFile(ctx context.Context, restoreFileReqDTO repoDTO.RestoreFileReqDTO) error
	FindAllFilesByOwner(ctx context.Context, findAllFilesByOwnerReqDTO repoDTO.FindAllFilesByOwnerReqDTO) (repoDTO.FindAllFilesByOwnerRespDTO, error)
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
