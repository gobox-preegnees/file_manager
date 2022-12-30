package postgresql

import (
	"context"

	repoDTO "github.com/gobox-preegnees/file_manager/internal/adapters/repo"
	// usecase "github.com/gobox-preegnees/file_manager/internal/domain/usecase"
	state "github.com/gobox-preegnees/file_manager/pkg/state"

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

// /////////////////////////////////////////////////////////////
func (p *postgresql) SaveOwner(ctx context.Context, username, folder string) (int, error) {

	sql :=
		`
		INSERT INTO owners (username, folder)
		VALUES ($1, $2)
		RETURNING owner_id
		`
	var id int
	err := p.conn.QueryRow(ctx, sql, username, folder).Scan(&id)
	return id, err
}

func (p *postgresql) RenameOwner(ctx context.Context, ownerID int, newFolderName string) error {

	sql :=
		`
		UPDATE owners
		SET folder=$1
		WHERE owner_id=$2
		RETURNING owner_id
		`

	_, err := p.conn.Exec(ctx, sql, newFolderName, ownerID)
	return err
}

func (p *postgresql) DeleteOwner(ctx context.Context, ownerID int) error {

	sql :=
		`
		UPDATE owners
		SET removed=true
		WHERE owner_id=$1
		`

	_, err := p.conn.Exec(ctx, sql, ownerID)
	return err
}

func (p *postgresql) RestoreOwner(ctx context.Context, ownerID string) error {

	sql :=
		`
		UPDATE owners
		SET removed=false
		WHERE owner_id=$1
		`
	_, err := p.conn.Exec(ctx, sql, ownerID)
	return err
}

func (p *postgresql) SaveFile(ctx context.Context, saveFileReqDTO repoDTO.SaveFileReqDTO) (int, error) {

	saveFile := func(state int) (int, error) {
		sql :=
			`
		INSERT INTO 
		files (client, file_name, mod_time, size_file, hash_sum, state, owner_id)
		VALUES (
			$1, $2, $3, $4, $5, $6
			(
				SELECT owner_id
				FROM owners
				WHERE removed=false
					AND folder=$7 
					AND username=$8
			)
		)
		RETURNING file_id
        `

		var id int
		err := p.conn.QueryRow(
			ctx, sql,
			saveFileReqDTO.Client, saveFileReqDTO.FileName,
			saveFileReqDTO.ModTime, saveFileReqDTO.SizeFile,
			saveFileReqDTO.HashSum, state,
			saveFileReqDTO.Folder, saveFileReqDTO.Username,
		).Scan(&id)
		return id, err
	}

	var err error
	var id int
	if saveFileReqDTO.HashSum == "" {
		id, err = saveFile(state.Folder)
	} else {
		id, err = saveFile(state.Created)
	}
	return id, err
}

func (p *postgresql) SetState(ctx context.Context, setStateReqDTO repoDTO.SetStateReqDTO) error {

	sql :=
		`
        UPDATE files
        SET state=$1, virtual_name=$2
        WHERE file_name=$3
			AND hash_sum=$4
			AND mod_time=$5
			AND size_file=$6
			AND owner_name=(
							SELECT owner_id
							FROM owners
							WHERE removed=false
								AND folder=$7
								AND username=$8
							)
        `

	_, err := p.conn.Exec(
		ctx, sql,
		setStateReqDTO.State, setStateReqDTO.VirtualName,
		setStateReqDTO.FileName, setStateReqDTO.HashSum,
		setStateReqDTO.ModTime, setStateReqDTO.SizeFile,
		setStateReqDTO.Folder, setStateReqDTO.Username,
	)
	return err
}

func (p *postgresql) RenameFile(ctx context.Context, renameFileReqDTO repoDTO.RenameFileReqDTO) error {

	sql :=
		`
			UPDATE files
			SET file_name=REPLACE(file_name, $1, $2), client=$3
			WHERE file_name LIKE "$4%"
				AND owner_id=(
							SELECT owner_id
							FROM owners
							WHERE removed=false
								AND folder=$5
								AND username=$6
							)
			`
	_, err := p.conn.Exec(
		ctx, sql,
		renameFileReqDTO.OldName, renameFileReqDTO.NewName,
		renameFileReqDTO.Client, renameFileReqDTO.OldName,
		renameFileReqDTO.Folder, renameFileReqDTO.Username,
	)
	return err
}

func (p *postgresql) DeleteFile(ctx context.Context, deleteFileReqDTO repoDTO.DeleteFileReqDTO) error {

	sql :=
		`
			UPDATE files
			SET removed=true, client=$1
			WHERE file_name LIKE "$2%"
				AND owner_id=(
							SELECT owner_id
							FROM owners
							WHERE removed=false
								AND folder=$3
								AND username=$4
							)
			RETURNING file_id
			`
	_, err := p.conn.Exec(
		ctx, sql,
		deleteFileReqDTO.Client, deleteFileReqDTO.FileName,
		deleteFileReqDTO.Folder, deleteFileReqDTO.Username,
	)

	return err
}

// RestoreFile. Restoring one file if it has a hash,
// or restores many files and folders if there is no
// hash (the folder does not have a hash, the folder contains files).
// You can restore only through your personal account (Web).
func (p *postgresql) RestoreFile(ctx context.Context, restoreFileReqDTO repoDTO.RestoreFileReqDTO) error {

	sql :=
		`
			UPDATE files
			SET removed=false, client=$1
			WHERE file_name LIKE "$2%"
				AND owner_id=(
							SELECT owner_id
							FROM owners
							WHERE removed=false
								AND folder=$3
								AND username=$4
							)
			RETURNING file_id
			`
	_, err := p.conn.Exec(
		ctx, sql,
		restoreFileReqDTO.Username, restoreFileReqDTO.FileName,
		restoreFileReqDTO.Folder, restoreFileReqDTO.Username,
	)
	return err
}

// var _ usecase.FileRepo = (*postgresql)(nil)
