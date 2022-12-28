package repo

import "github.com/gobox-preegnees/file_manager/internal/domain/entity"

type FileDTO struct {
	Client string
	entity.File
	State int
}