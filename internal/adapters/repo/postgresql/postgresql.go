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
		INNER JOIN files ON folders.folder_id = files.folder_id
		WHERE folders.removed = FALSE 
			AND users.removed = FALSE 
			AND users.username = $1 
			AND folders.folder = $2
		`
	rows, err := p.conn.Query(ctx, sql, username, folder)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var files []repoDTO.FileDTO = make([]repoDTO.FileDTO, 0)
	for rows.Next() {
		var file repoDTO.FileDTO = repoDTO.FileDTO{}
		err := rows.Scan(
			&file.Client,
			&file.FileName,
			&file.HashSum,
			&file.SizeFile,
			&file.ModTime,
			&file.VirtualName,
			&file.State,
			&file.Removed,
		)
		if err != nil {
			return nil, err
		}
		files = append(files, file)
	}
	return files, nil
}

func (p *postgresql) FindOneByPath(ctx context.Context, username string, folder string, fileName string) (repoDTO.FileDTO, error) {

	sql :=
		`
		SELECT client, file_name, hash_sum, size_file, mod_time, virtual_name, state, files.removed 
		FROM users
		INNER JOIN folders ON folders.user_id = users.user_id
		INNER JOIN files ON folders.folder_id = files.folder_id
		WHERE folders.removed = FALSE 
			AND users.removed = FALSE
			AND users.username = $1 
			AND folders.folder = $2
			AND file_name = $3
		`
	var file repoDTO.FileDTO = repoDTO.FileDTO{}
	err := p.conn.QueryRow(ctx, sql, username, folder, fileName).Scan(
		&file.Client,
		&file.FileName,
		&file.HashSum,
		&file.SizeFile,
		&file.ModTime,
		&file.VirtualName,
		&file.State,
		&file.Removed,
	)
	if err != nil {
		return repoDTO.FileDTO{}, err
	}
	return file, nil
}

func (p *postgresql) SaveOne(ctx context.Context, username, folder string, file repoDTO.FileDTO) error {

	sql :=
		`
		INSERT INTO 
		files (client,file_name,mod_time,size_file,hash_sum,folder_id)
		VALUES (
			$1,$2,$3,$4,$5,
			(
				SELECT folder_id
				FROM folders
				WHERE removed=false
					AND folder=$6 
					AND user_id=(
								SELECT user_id 
								FROM users
								WHERE username=$7
									AND removed=false
								)
			)
		);
		`
	_, err := p.conn.Exec(ctx, sql,
		file.Client, file.FileName, file.ModTime, file.SizeFile, file.HashSum, folder, username,
	)
	return err
}

func (p *postgresql) UpdateState(ctx context.Context, username, folder, client, fileName, hashSum, virtualName string, state int) error {

	sql :=
		`
		UPDATE files 
		SET state=$1, client=$2, virtual_name=$3
		WHERE hash_sum=$4 
			AND file_name=$5
			AND folder_id=(
						SELECT folder_id
						FROM folders
						WHERE removed=false
							AND folder=$6
							AND user_id=(
										SELECT user_id
										FROM users
										WHERE username=$7
											AND removed=false
										)
						)
		`
	_, err := p.conn.Exec(ctx, sql, state, client, virtualName, hashSum, fileName, folder, username)
	return err
}

func (p *postgresql) UpdateFileName(ctx context.Context, username, folder, client, oldfileName, newfileName, hash string) error {
	var err error
	if hash == "" {
		sql :=
			`
			UPDATE files
			SET file_name=REPLACE(file_name, $1, $2), client=$3
			WHERE folder_id=(
							SELECT folder_id
							FROM folders
							WHERE removed=false
								AND folder=$4
								AND user_id=(
											SELECT user_id
											FROM users
											WHERE username=$5
												AND removed=false
											)
							)
			`
		_, err = p.conn.Exec(ctx, sql, oldfileName, newfileName, client, folder, username)
	} else {
		sql :=
			`
			UPDATE files
			SET file_name=REPLACE(file_name, $1, $2), client=$3
			WHERE hash_sum=$4
				AND folder_id=(
							SELECT folder_id
							FROM folders
							WHERE folder=$5
								AND user_id=(
											SELECT user_id
											FROM users
											WHERE username=$6
											)
							)
			`
		_, err = p.conn.Exec(ctx, sql, oldfileName, newfileName, client, hash, folder, username)
	}
	return err
}

func (p *postgresql) DeleteFile(ctx context.Context, username, folder, client, fileName, hash string) error {
	var err error
	if hash == "" {
		sql :=
			`
		UPDATE files
		SET removed=true, client=$1
		WHERE file_name LIKE "$2%"
			AND folder_id=(
						SELECT folder_id
						FROM folders
						WHERE folder=$3
							AND user_id=(
										SELECT user_id
										FROM users
										WHERE username=$4
										)
						)
		`
		_, err = p.conn.Exec(ctx, sql, client, fileName, folder, username)
	} else {
		sql :=
			`
		UPDATE files
		SET removed=true, client=$1
		WHERE file_name LIKE "$2%"
			AND hash_sum=$3
			AND folder_id=(
						SELECT folder_id
						FROM folders
						WHERE folder=$4
							AND user_id=(
										SELECT user_id
										FROM users
										WHERE username=$5
										)
						)
		`
		_, err = p.conn.Exec(ctx, sql, client, fileName, hash, folder, username)
	}

	return err
}

func (p *postgresql) CreateUser(ctx context.Context, username string) error {
	sql :=
		`
		INSERT INTO users (username)
		VALUES (
			$1
		)
		`
	_, err := p.conn.Exec(ctx, sql, username)
	return err
}

func (p *postgresql) DeleteUser(ctx context.Context, username string) error {
	sql :=
		`
		UPDATE users
		SET removed=true
		WHERE username=$1
		`
	_, err := p.conn.Exec(ctx, sql, username)
	return err
}

func (p *postgresql) CreateFolder(ctx context.Context, username, folder string) error {
	sql :=
		`
		INSERT INTO folders (folder, user_id)
		VALUES (
			$1,
			(
				SELECT user_id 
				FROM users 
				WHERE username=$2 
					AND removed=true
			)
		)
		`
	_, err := p.conn.Exec(ctx, sql, username)
	return err
}

func (p *postgresql) DeleteFolder(ctx context.Context, username, folder string) error {
	sql :=
		`
		UPDATE folders
		SET removed=true
		WHERE folder=$1
			AND folder_id=(
						SELECT user_id 
						FROM users
						WHERE username=$2
						)
		`
	_, err := p.conn.Exec(ctx, sql, folder, username)
	return err
}

var _ usecase.FileRepo = (*postgresql)(nil)
