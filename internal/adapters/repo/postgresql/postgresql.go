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

func (p *postgresql) FindAll(ctx context.Context, username string, folder string) ([]repoDTO.FileDTO, error) {

	sql :=
		`
		SELECT client, file_name, hash_sum, size_file, mod_time, virtual_name, state, files.removed 
		FROM users
		INNER JOIN folders ON folders.user_id = users.user_id
		INNER JOIN files ON users.user_id = files.user_id
		WHERE folders.removed = FALSE 
			AND users.username = $1 
			AND folders.folder = $2
		`
	rows, err := p.conn.Query(ctx, sql, username, folder)
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
			&item.SizeFile,
			&item.ModTime,
			&item.VirtualName,
			&item.State,
			&item.Removed,
		)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, nil
}

func (p *postgresql) FindOneByPath(ctx context.Context, username string, folder string, fileName string) (repoDTO.FileDTO, error) {
	
	sql :=
		`
		SELECT client, file_name, hash_sum, size_file, mod_time, virtual_name, state, files.removed 
		FROM users
		INNER JOIN folders ON folders.user_id = users.user_id
		INNER JOIN files ON users.user_id = files.user_id
		WHERE folders.removed = FALSE 
			AND users.username = $1 
			AND folders.folder = $2
			AND file_name = $3
		`
    item := repoDTO.FileDTO{}
	err := p.conn.QueryRow(ctx, sql, username, folder, fileName).Scan(
		&item.Client,
		&item.FileName,
		&item.HashSum,
		&item.SizeFile,
		&item.ModTime,
		&item.VirtualName,
		&item.State,
        &item.Removed,
	)
	if err != nil {
		return repoDTO.FileDTO{}, err
	}
	return item, nil
}

func (*postgresql) Save(ctx context.Context, username, folder string, file repoDTO.FileDTO) error {

	sql :=
        `
		
		`
	return nil
}

func (*postgresql) UpdateState(ctx context.Context, username, folder, client, fileName string) (error) {
	return nil
}

func (*postgresql) UpdateFileName(ctx context.Context, username, folder, client, oldfileName, newfileName, hash string) (error) {
	return nil
}

func (*postgresql) DeleteFile(ctx context.Context, username, folder, client, fileName, hash string) (error) {
	return nil
}



var _ usecase.FileRepo = (*postgresql)(nil)
