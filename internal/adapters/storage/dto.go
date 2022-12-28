package postgresql

import (
	entity "github.com/gobox-preegnees/file_manager/internal/domain/entity"
)

type UpdateOneReqDTO struct {
	entity.Identifier
	OldPath string
	NewPath string
}
type UpdateAllReqDTO struct {
	entity.Identifier
	OldPath string
	NewPath string
}
type CreateReqDTO struct {
	entity.Identifier
	entity.File
}
type FindOneReqDTO struct {
	entity.Identifier
}
type FindOneRespDTO struct {
	File entity.File
}
type DeleteOneReqDTO struct {
	entity.Identifier
	Path string
}
type DeleteAllReqDTO struct {
	entity.Identifier
	Path string
}
type FindAllReqDTO struct {
	entity.Identifier
	Path string
}
type FindAllRespDTO struct {
	Files []entity.File
}
type CreateRespDTO struct {
	ID int
}