package disk

import (
	"context"
	"os"

	// dtoStorage "github.com/gobox-preegnees/file_manager/v1/internal/adapters/storage"
	entity "github.com/gobox-preegnees/file_manager/v1/internal/domain/entity"
)

type diskStorage struct {
	file *os.File
}

func New(file *os.File) diskStorage {
	return diskStorage{
		file: file,
	}
}

func (d diskStorage) Create(ctx context.Context, file entity.File) error {
	if _, err := d.file.Write([]byte(file.Path)); err != nil {
		return err
	}
	return nil
}

func (d diskStorage) GetByPath(ctx context.Context, path entity.Path) (file entity.File, err error) {
	return entity.File{}, nil
}
func (d diskStorage) Update(ctx context.Context, file entity.File) (id int64, err error) {
	return -1, nil
}
func (d diskStorage)  DeleteByPath(ctx context.Context, path entity.Path) (id int64, err error) {
	return -1, nil
}
