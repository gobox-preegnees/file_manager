package postgresql

import (
	"context"
	"errors"
	"fmt"

	repoDTO "github.com/gobox-preegnees/file_manager/internal/adapters/repo"
	usecase "github.com/gobox-preegnees/file_manager/internal/domain/usecase"
	state "github.com/gobox-preegnees/file_manager/pkg/state"

	"github.com/jackc/pgx/v5"
)

var (
	ErrInvalidOwner = errors.New("Invalid Owner On FindAllByOwner")
)

// postgresql. Implementation of the usecase package interface
type postgresql struct {
	conn *pgx.Conn
}

// New. Create a new Postgres
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

// Save Owner. Stores UserName + Folder . This method will participate in a distributed transaction
func (p *postgresql) SaveOwner(ctx context.Context, saveOwnerDTO repoDTO.SaveOwnerDTO) (int, error) {

	sql :=
		`
		INSERT INTO owners (username, folder)
		VALUES ($1, $2)
		RETURNING owner_id
		`
	var id int
	err := p.conn.QueryRow(ctx, sql, saveOwnerDTO.Username, saveOwnerDTO.Folder).Scan(&id)
	return id, err
}

// Rename Owner. Renames Folder. This method will participate in a distributed transaction
func (p *postgresql) RenameOwner(ctx context.Context, renameOwnerDTO repoDTO.RenameOwnerDTO) error {

	sql :=
		`
		UPDATE owners
		SET folder=$1
		WHERE owner_id=$2
		RETURNING owner_id
		`
	_, err := p.conn.Exec(ctx, sql, renameOwnerDTO.NewName, renameOwnerDTO.OwnerID)
	return err
}

// Delete Owner. Removes the owner (the user will not be able to restore). This method will participate in a distributed transaction
func (p *postgresql) DeleteOwner(ctx context.Context, deleteOwnerDTO repoDTO.DeleteOwnerDTO) error {

	sql :=
		`
		UPDATE owners
		SET removed=true
		WHERE owner_id=$1
		`
	_, err := p.conn.Exec(ctx, sql, deleteOwnerDTO.OwnerID)
	return err
}

// FindAllOwners. Gets all Folders (Folders that are synchronized, virtual folders, not real ones) by username
func (p *postgresql) FindAllOwners(ctx context.Context, findAllOwnersReqDTO repoDTO.FindAllOwnersReqDTO) (repoDTO.FindAllOwnersRespDTO, error) {

	sql :=
		`
        SELECT owner_id, username, folder
		FROM owners
        WHERE username=$1
        `
	rows, err := p.conn.Query(ctx, sql, findAllOwnersReqDTO.Username)
	if err != nil {
		return repoDTO.FindAllOwnersRespDTO{}, err
	}
	defer rows.Close()

	owners := make([]repoDTO.Owner, 0)
	for rows.Next() {
		owner := repoDTO.Owner{}
		err := rows.Scan(&owner.OwnerId, &owner.Username, &owner.Folder)
		if err != nil {
			return repoDTO.FindAllOwnersRespDTO{}, err
		}
		owners = append(owners, owner)
	}
	return repoDTO.FindAllOwnersRespDTO{
		Owners: owners,
	}, nil
}

// Save file. Saves the file, for the file the status is = 100, 
// for the folder = 300 (since the folder has no size, hash and does not need to be uploaded to the server)
func (p *postgresql) SaveFile(ctx context.Context, saveFileReqDTO repoDTO.SaveFileReqDTO) (int, error) {

	saveFile := func(state int) (int, error) {
		sql :=
			`
			INSERT INTO 
			files (client, file_name, mod_time, size_file, hash_sum, state, owner_id)
			VALUES (
				$1, $2, $3, $4, $5, $6,
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

// SetState. Sets the state of the file at the current moment. 
// This only applies to files. When a file is uploaded to the server, it saves its state:
// 1. Created: when saving to the database (before the actual loading), the file has such a state, code = 100.
// 2. Uploading: when the file starts uploading to the server, it has code = 200.
// 3. Uploaded: when the file is definitely uploaded to the server, it has code = 300.
// 4. For any error during the upload process, the file code becomes 400.
// VirtualName - file name on the disk (randomly generated on save).
func (p *postgresql) SetState(ctx context.Context, setStateReqDTO repoDTO.SetStateReqDTO) error {

	sql :=
		`
        UPDATE files
        SET state=$1, virtual_name=$2
        WHERE file_name=$3
			AND hash_sum=$4
			AND mod_time=$5
			AND size_file=$6
			AND owner_id=(
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

// RenameFile. Renames files by file name, if this is a folder name, 
// then all files and folders in it will be renamed
func (p *postgresql) RenameFile(ctx context.Context, renameFileReqDTO repoDTO.RenameFileReqDTO) error {

	sql :=
		`
			UPDATE files
			SET file_name=REPLACE(file_name, $1, $2), client=$3
			WHERE file_name LIKE $4
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
		renameFileReqDTO.Client, fmt.Sprintf("%s%s%s", "", renameFileReqDTO.OldName, "%"),
		renameFileReqDTO.Folder, renameFileReqDTO.Username,
	)
	return err
}

// DeleteFile. Deletes files by file name, if this is a folder name, 
// then all files and folders in it will be deleted
func (p *postgresql) DeleteFile(ctx context.Context, deleteFileReqDTO repoDTO.DeleteFileReqDTO) error {

	sql :=
		`
			UPDATE files
			SET removed=true, client=$1
			WHERE file_name LIKE $2
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
		deleteFileReqDTO.Client, fmt.Sprintf("%s%s%s", "", deleteFileReqDTO.FileName, "%"),
		deleteFileReqDTO.Folder, deleteFileReqDTO.Username,
	)
	return err
}

// RestoreFile. Restoring one file if it has a hash,
// or restores many files and folders if there is no
// hash (there is no hash in the folder, the folder contains files and they will be restored).
// You can restore only through your personal account (Web).
func (p *postgresql) RestoreFile(ctx context.Context, restoreFileReqDTO repoDTO.RestoreFileReqDTO) error {

	sql :=
		`
		UPDATE files
		SET removed=false, client=$1
		WHERE file_name LIKE $2
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
		restoreFileReqDTO.Username, fmt.Sprintf("%s%s%s", "", restoreFileReqDTO.FileName, "%"),
		restoreFileReqDTO.Folder, restoreFileReqDTO.Username,
	)
	return err
}

// FindAllFilesByOwner. Gets all files by owner.
// You can get both by owner_id or by username + folder.
func (p *postgresql) FindAllFilesByOwner(ctx context.Context, findAllFilesByOwnerReqDTO repoDTO.FindAllFilesByOwnerReqDTO) (repoDTO.FindAllFilesByOwnerRespDTO, error) {

	sql :=
		`
		SELECT file_id, files.removed, virtual_name, state, hash_sum, file_name, size_file, owner_id, mod_time, client, username, folder
		FROM owners
		INNER JOIN files ON files.owner_id=owners.owner_id
		WHERE owners.removed=false 
		`
	// TODO: add without ByOwner -> ..., search by any field
	var err error
	var rows pgx.Rows
	if findAllFilesByOwnerReqDTO.OwnerId != 0 {
		sql = sql + " AND owner_id=$1"
		rows, err = p.conn.Query(ctx, sql, findAllFilesByOwnerReqDTO.OwnerId)
	} else if findAllFilesByOwnerReqDTO.Folder != "" && findAllFilesByOwnerReqDTO.Username != "" {
		sql = sql + " AND folder=$1 AND username=$2"
		rows, err = p.conn.Query(ctx, sql, findAllFilesByOwnerReqDTO.OwnerId)
	} else {
		err = ErrInvalidOwner
	}
	if err != nil {
		return repoDTO.FindAllFilesByOwnerRespDTO{}, err
	}
	defer rows.Close()

	var files = make([]repoDTO.FullFile, 0)
	for rows.Next() {
		var file repoDTO.FullFile = repoDTO.FullFile{}
		err := rows.Scan(
			&file.File_id,
			&file.Removed,
			&file.VirtualName,
			&file.State,
			&file.HashSum,
			&file.FileName,
			&file.SizeFile,
			&file.OwnerId,
			&file.ModTime,
			&file.Client,
			&file.Username,
			&file.Folder,
		)
		if err != nil {
			return repoDTO.FindAllFilesByOwnerRespDTO{}, err
		}
		files = append(files, file)
	}
	return repoDTO.FindAllFilesByOwnerRespDTO{
		Files: files,
	}, nil
}

var _ usecase.FileRepo = (*postgresql)(nil)