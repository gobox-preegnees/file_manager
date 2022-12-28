package postgresql

import (
	"context"

	storageDTO "github.com/gobox-preegnees/file_manager/internal/adapters/storage"
	usecase "github.com/gobox-preegnees/file_manager/internal/domain/usecase"

	"github.com/jackc/pgx/v5"
)

type postgresql struct{
	conn *pgx.Conn
}

func New(ctx context.Context, url string) (*postgresql, error) {

	conn, err := pgx.Connect(ctx, url)
	if err!= nil {
        return nil, err
    }
	if err := conn.Ping(ctx); err != nil {
		return nil, err
	}

	return &postgresql{conn: conn}, nil
}

func (p *postgresql) CreateOne(ctx context.Context, createReqDTO storageDTO.CreateReqDTO) error {
	return nil
}

func (*postgresql) DeleteAll(ctx context.Context, deleteAllReqDTO storageDTO.DeleteAllReqDTO) error {
	return nil
}

func (*postgresql) DeleteOne(ctx context.Context, deleteOneReqDTO storageDTO.DeleteOneReqDTO) error {
	return nil
}

func (*postgresql) FindAll(ctx context.Context, findAllReqDTO storageDTO.FindAllReqDTO) (FindAllRespDTO storageDTO.FindAllRespDTO, err error) {
	return storageDTO.FindAllRespDTO{}, nil
}

func (*postgresql) FindOneByPath(ctx context.Context, findOneReqDTO storageDTO.FindOneReqDTO) (FindOneRespDTO storageDTO.FindOneRespDTO, err error) {
	return storageDTO.FindOneRespDTO{}, nil
}

func (*postgresql) UpdateAll(ctx context.Context, updateAllReqDTO storageDTO.UpdateAllReqDTO) error {
	return nil
}

func (*postgresql) UpdateOne(ctx context.Context, updateOneReqDTO storageDTO.UpdateOneReqDTO) error {
	return nil
}

var _ usecase.IFileStorage = (*postgresql)(nil)
