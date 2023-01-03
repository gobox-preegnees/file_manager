package domain

import (
	"github.com/gobox-preegnees/file_manager/internal/domain/entity"
)

type GetFilesReqDTO struct {
	Identifier      entity.Identifier
	OwnerId, FileId int
}

type SaveFileReqDTO struct {
	Identifier entity.Identifier
	File       entity.File
	Client     string
}

type RenameFileReqDTO struct {
	Identifier entity.Identifier
	OldFilName, NewFileName, Client string
}

type DeleteFileReqDTO struct {
	Identifier entity.Identifier
	Client, FileName string
}

type SetStateReqDTO struct {
	entity.State
}

type CreateOwnerReqDTO struct {
	Owner entity.Owner
}

type DeleteOwnerReqDTO struct {
	Id int
}