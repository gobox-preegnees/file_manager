package usecase

import (
	"context"

	daoDTO "github.com/gobox-preegnees/file_manager/internal/adapters/dao"
	grpcController "github.com/gobox-preegnees/file_manager/internal/controller/grpc"
	entity "github.com/gobox-preegnees/file_manager/internal/domain/entity"
	"github.com/sirupsen/logrus"
)

//go:generate mockgen -destination=../../mocks/domain/usecase/dao/file/file.go -package=usecase_dao_file -source=file.go
type IDaoFile interface {
	SaveFile(ctx context.Context, saveFileReqDTO daoDTO.SaveFileReqDTO) (int, error)
	RenameFile(ctx context.Context, renameFileReqDTO daoDTO.RenameFileReqDTO) error
	DeleteFile(ctx context.Context, deleteFileReqDTO daoDTO.DeleteFileReqDTO) error
	FindAllFilesByOwnerOrFileId(ctx context.Context, findAllFilesByOwnerReqDTO daoDTO.FindAllFilesByOwnerOrFileIdReqDTO) (daoDTO.FindAllFilesByOwnerOrFileIdRespDTO, error)
}

type fileUsecase struct {
	log      *logrus.Logger
	fileRepo IDaoFile
}

func NewFileUsecase(log *logrus.Logger, fileRepo IDaoFile) *fileUsecase {

	return &fileUsecase{
		log:      log,
		fileRepo: fileRepo,
	}
}

func (f *fileUsecase) GetFiles(ctx context.Context, identifier entity.Identifier, ownerId, fileId int) ([]daoDTO.FullFile, error) {

	filesDTO, err := f.fileRepo.FindAllFilesByOwnerOrFileId(ctx, daoDTO.FindAllFilesByOwnerOrFileIdReqDTO{
		Owner: daoDTO.Owner{
			Identifier: daoDTO.Identifier{
				Username: identifier.Username,
				Folder:   identifier.Folder,
			},
			OwnerId: ownerId,
		},
		FileId: fileId,
	})
	if err != nil {
		return nil, err
	}
	return filesDTO.Files, nil
}

func (f *fileUsecase) SaveFile(ctx context.Context, identifier entity.Identifier, file entity.File, client string) (int, error) {

	id, err := f.fileRepo.SaveFile(ctx, daoDTO.SaveFileReqDTO{
		Identifier: daoDTO.Identifier{
			Username: identifier.Username,
			Folder:   identifier.Folder,
		},
		File: daoDTO.File{
			FileName: file.FileName,
			HashSum:  file.HashSum,
			SizeFile: file.SizeFile,
			ModTime:  file.ModTime,
		},
		Client: client,
	})
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (f *fileUsecase) RenameFile(ctx context.Context, identifier entity.Identifier, oldFilName, newFileName, client string) error {

	err := f.fileRepo.RenameFile(ctx, daoDTO.RenameFileReqDTO{
		Identifier: daoDTO.Identifier{
			Username: identifier.Username,
			Folder:   identifier.Folder,
		},
		Client:  client,
		OldName: oldFilName,
		NewName: newFileName,
	})
	return err
}

func (f *fileUsecase) DeleteFile(ctx context.Context, identifier entity.Identifier, client, fileName string) error {

	err := f.fileRepo.DeleteFile(ctx, daoDTO.DeleteFileReqDTO{
		Identifier: daoDTO.Identifier{
			Username: identifier.Username,
			Folder:   identifier.Folder,
		},
		Client:   client,
		FileName: fileName,
	})
	return err
}

var _ grpcController.IFileUsecase = (*fileUsecase)(nil)
