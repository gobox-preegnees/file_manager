package file_usecase

import (
	"context"

	entity "github.com/gobox-preegnees/file_manager/v1/internal/domain/entity"
)

type IFileService interface {
	SaveFile(ctx context.Context, file entity.File) error
}

type IFolderService interface {}

type FileUsecase struct {
	fileService IFileService 
	folderService IFolderService
}

func New(fileService IFileService, folderService IFolderService) FileUsecase {
	return FileUsecase{
		fileService: fileService,
		folderService: folderService,
	}
}

func (f FileUsecase) SaveFile(ctx context.Context, file SaveFileDTO) error {
	return f.fileService.SaveFile(ctx, entity.File{
		Identifier: entity.Identifier{
			Username: entity.Username(file.FullFile.Username),
			FolderID: entity.FolderID(file.FullFile.FolderID),
			ClientID: entity.ClientID(file.FullFile.ClientID),
		},
		Path: entity.Path(file.Path),
		Hash: entity.Hash(file.Hash),
		Size: entity.Size(file.Size),
		ModTime: entity.ModTime(file.ModTime),
		VirtualName: entity.VirtualName(file.VirtualName),
	})
}