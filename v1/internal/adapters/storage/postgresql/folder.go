package postgresql

import (
	"context"

	"github.com/jackc/pgx/v5"

	entity "github.com/gobox-preegnees/file_manager/v1/internal/domain/entity"
	// services "github.com/gobox-preegnees/file_manager/v1/internal/domain/services"
)

type folderStorage struct {
	conn *pgx.Conn
}

// var _ services.IFolderStorage = (*folderStorage)(nil)

func NewFolderStorage(conn *pgx.Conn) *folderStorage {
	return &folderStorage{
		conn: conn,
	}
}

func (f *folderStorage) Ping(ctx context.Context) error {
	return f.conn.Ping(ctx)
}

func (f *folderStorage) GetByPath(ctx context.Context, path entity.Path) (file entity.Folder, err error) {
	panic("unimplemented")
}

func (f *folderStorage) GetByModTime(ctx context.Context, modTime entity.ModTime) (file entity.Folder, err error) {
	panic("unimplemented")
}

func (f *folderStorage) Create(ctx context.Context, folder entity.Folder) (id int64, err error) {
	panic("unimplemented")
}

func (f *folderStorage) Update(ctx context.Context, folder entity.Folder) (id int64, err error) {
	panic("unimplemented")
}

func (f *folderStorage) DeleteByPath(ctx context.Context, path entity.Path) (id int64, err error) {
	panic("unimplemented")
}

func (f *folderStorage) DeleteByModTime(ctx context.Context, modTime entity.ModTime) (id int64, err error) {
	panic("unimplemented")
}
