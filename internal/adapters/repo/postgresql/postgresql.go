package postgresql

import (
	"context"

	repoDTO "github.com/gobox-preegnees/file_manager/internal/adapters/repo"
	usecase "github.com/gobox-preegnees/file_manager/internal/domain/usecase"

	"github.com/jackc/pgx/v5"
)

type postgresql struct {
	conn *pgx.Conn
}

func New(ctx context.Context, url string) (*postgresql, error) {

	conn, err := pgx.Connect(ctx, url)
	if err != nil {
		return nil, err
	}
	if err := conn.Ping(ctx); err != nil {
		return nil, err
	}

	return &postgresql{conn: conn}, nil
}

func (p *postgresql) FindAll(ctx context.Context, username string, folderID string) ([]repoDTO.FileDTO, error) {

	sql :=
	`
	SELECT client, file_name, hash_sum, size_file, mod_time, virtual_name, state 
	FROM users
	INNER JOIN folders ON folders.user_id = users.user_id
	INNER JOIN files ON users.user_id = files.user_id
	INNER JOIN states ON files.file_id = states.file_id
	`
	rows, err := p.conn.Query(ctx, sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

    var items []repoDTO.FileDTO = make([]repoDTO.FileDTO, 0)
	for rows.Next() {
		var item repoDTO.FileDTO = repoDTO.FileDTO{}
		err := rows.Scan(
			&item.Client,
            &item.FileName,
            &item.HashSum,
            &item.FileSize,
            &item.ModTime,
			&item.VirtualName,
			&item.State,
		)
	
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, nil
}

func (*postgresql) FindOneByPath(ctx context.Context, username string, folderID string, path string) (repoDTO.FileDTO, error) {
	return repoDTO.FileDTO{}, nil
}

func (*postgresql) Save(ctx context.Context, file repoDTO.FileDTO) error {
	return nil
}

var _ usecase.FileRepo = (*postgresql)(nil)
