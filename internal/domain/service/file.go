package service

import (
	"context"

	daoDTO "github.com/gobox-preegnees/file_manager/internal/adapters/dao"
	dtoService "github.com/gobox-preegnees/file_manager/internal/domain"
	grpcController "github.com/gobox-preegnees/file_manager/internal/controller/grpc"
	"github.com/sirupsen/logrus"
)

//go:generate mockgen -destination=../../mocks/domain/usecase/dao/file/file.go -package=usecase_dao_file -source=file.go
type IDaoFile interface {
	SaveFile(ctx context.Context, saveFileReqDTO daoDTO.SaveFileReqDTO) (int, error)
	RenameFile(ctx context.Context, renameFileReqDTO daoDTO.RenameFileReqDTO) error
	DeleteFile(ctx context.Context, deleteFileReqDTO daoDTO.DeleteFileReqDTO) error
	FindAllFilesByOwnerOrFileId(ctx context.Context, findAllFilesByOwnerReqDTO daoDTO.FindAllFilesByOwnerOrFileIdReqDTO) (daoDTO.FindAllFilesByOwnerOrFileIdRespDTO, error)
}

type fileService struct {
	log      *logrus.Logger
	daoFile IDaoFile
}

func NewFileUsecase(log *logrus.Logger, daoFile IDaoFile) *fileService {

	return &fileService{
		log:      log,
		daoFile: daoFile,
	}
}

func (f fileService) GetFiles(ctx context.Context, getFilesReqDTO dtoService.GetFilesReqDTO) ([]daoDTO.FullFile, error) {

	filesDTO, err := f.daoFile.FindAllFilesByOwnerOrFileId(ctx, daoDTO.FindAllFilesByOwnerOrFileIdReqDTO{
		Owner: daoDTO.Owner{
			Identifier: daoDTO.Identifier{
				Username: getFilesReqDTO.Identifier.Username,
				Folder:   getFilesReqDTO.Identifier.Folder,
			},
			OwnerId: getFilesReqDTO.OwnerId,
		},
		FileId: getFilesReqDTO.FileId,
	})
	if err != nil {
		return nil, err
	}
	return filesDTO.Files, nil
}

func (f fileService) SaveFile(ctx context.Context, saveFileReqDTO dtoService.SaveFileReqDTO) (int, error) {

	id, err := f.daoFile.SaveFile(ctx, daoDTO.SaveFileReqDTO{
		Identifier: daoDTO.Identifier{
			Username: saveFileReqDTO.Identifier.Username,
			Folder:   saveFileReqDTO.Identifier.Folder,
		},
		File: daoDTO.File{
			FileName: saveFileReqDTO.File.FileName,
			HashSum:  saveFileReqDTO.File.HashSum,
			SizeFile: saveFileReqDTO.File.SizeFile,
			ModTime:  saveFileReqDTO.File.ModTime,
		},
		Client: saveFileReqDTO.Client,
	})
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (f fileService) RenameFile(ctx context.Context, renameFileReqDTO dtoService.RenameFileReqDTO) error {

	err := f.daoFile.RenameFile(ctx, daoDTO.RenameFileReqDTO{
		Identifier: daoDTO.Identifier{
			Username: renameFileReqDTO.Identifier.Username,
			Folder:   renameFileReqDTO.Identifier.Folder,
		},
		Client:  renameFileReqDTO.Client,
		OldName: renameFileReqDTO.OldFilName,
		NewName: renameFileReqDTO.NewFileName,
	})
	return err
}

func (f fileService) DeleteFile(ctx context.Context, deleteFileReqDTO dtoService.DeleteFileReqDTO) error {

	err := f.daoFile.DeleteFile(ctx, daoDTO.DeleteFileReqDTO{
		Identifier: daoDTO.Identifier{
			Username: deleteFileReqDTO.Identifier.Username,
			Folder:   deleteFileReqDTO.Identifier.Folder,
		},
		Client:   deleteFileReqDTO.Client,
		FileName: deleteFileReqDTO.FileName,
	})
	return err
}

var _ grpcController.IFileService = (*fileService)(nil)
