package postgresql

import (
	"context"

	"github.com/jackc/pgx/v5"

	entity "github.com/gobox-preegnees/file_manager/v1/internal/domain/entity"
	// services "github.com/gobox-preegnees/file_manager/v1/internal/domain/services"
)

type fileStorage struct {
	conn *pgx.Conn
}

// var _ services.IFileStorage = (*fileStorage)(nil)

func NewFileStorage(conn *pgx.Conn) *fileStorage {
	return &fileStorage{
		conn: conn,
	}
}

func (f *fileStorage) Ping(ctx context.Context) error {
	return f.conn.Ping(ctx)
}

func (f *fileStorage) Create(ctx context.Context, file entity.File) (id int64, err error) {
	panic("unimplemented")
}

func (f *fileStorage) GetByPath(ctx context.Context, path entity.Path) (file entity.File, err error) {
	panic("unimplemented")
}

func (f *fileStorage) Update(ctx context.Context, file entity.File) (id int64, err error) {
	panic("unimplemented")
}

func (f *fileStorage) DeleteByPath(ctx context.Context, path entity.Path) (id int64, err error) {
	panic("unimplemented")
}
