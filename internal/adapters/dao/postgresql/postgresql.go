package postgresql

import (
	"context"
	"errors"
	"fmt"

	daoDTO "github.com/gobox-preegnees/file_manager/internal/adapters/dao"
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
func (p *postgresql) SaveOwner(ctx context.Context, saveOwnerDTO daoDTO.SaveOwnerDTO) (int, error) {

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
func (p *postgresql) RenameOwner(ctx context.Context, renameOwnerDTO daoDTO.RenameOwnerDTO) error {

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
func (p *postgresql) DeleteOwner(ctx context.Context, deleteOwnerDTO daoDTO.DeleteOwnerDTO) error {

	// TODO: удалить нормально
	sql :=
		`
		DELETE FROM owners 
		WHERE owner_id=$1
		`
	_, err := p.conn.Exec(ctx, sql, deleteOwnerDTO.OwnerID)
	return err
}

// FindAllOwners. Gets all Folders (Folders that are synchronized, virtual folders, not real ones) by username
func (p *postgresql) FindAllOwners(ctx context.Context, findAllOwnersReqDTO daoDTO.FindAllOwnersReqDTO) (daoDTO.FindAllOwnersRespDTO, error) {

	sql :=
		`
        SELECT owner_id, username, folder
		FROM owners
        WHERE username=$1
        `
	rows, err := p.conn.Query(ctx, sql, findAllOwnersReqDTO.Username)
	if err != nil {
		return daoDTO.FindAllOwnersRespDTO{}, err
	}
	defer rows.Close()

	owners := make([]daoDTO.Owner, 0)
	for rows.Next() {
		owner := daoDTO.Owner{}
		err := rows.Scan(&owner.OwnerId, &owner.Username, &owner.Folder)
		if err != nil {
			return daoDTO.FindAllOwnersRespDTO{}, err
		}
		owners = append(owners, owner)
	}
	return daoDTO.FindAllOwnersRespDTO{
		Owners: owners,
	}, nil
}

// Save file. Saves the file, for the file the status is = 100,
// for the folder = 300 (since the folder has no size, hash and does not need to be uploaded to the server)
func (p *postgresql) SaveFile(ctx context.Context, saveFileReqDTO daoDTO.SaveFileReqDTO) (int, error) {

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
					WHERE folder=$7 
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
func (p *postgresql) SetState(ctx context.Context, setStateReqDTO daoDTO.SetStateReqDTO) error {

	sql :=
		`
        UPDATE files
        SET state=$1, virtual_name=$2
        WHERE file_name=$3
			AND state < $4
			AND owner_id=(
							SELECT owner_id
							FROM owners
							WHERE folder=$5
								AND username=$6
							)
        `

	_, err := p.conn.Exec(
		ctx, sql,
		setStateReqDTO.State, setStateReqDTO.VirtualName,
		setStateReqDTO.FileName, setStateReqDTO.State,
		setStateReqDTO.Folder, setStateReqDTO.Username,
	)
	return err
}

// RenameFile. Renames files by file name, if this is a folder name,
// then all files and folders in it will be renamed
func (p *postgresql) RenameFile(ctx context.Context, renameFileReqDTO daoDTO.RenameFileReqDTO) error {

	sql :=
		`
			UPDATE files
			SET file_name=REPLACE(file_name, $1, $2), client=$3
			WHERE file_name LIKE $4
				AND owner_id=(
							SELECT owner_id
							FROM owners
							WHERE folder=$5
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
func (p *postgresql) DeleteFile(ctx context.Context, deleteFileReqDTO daoDTO.DeleteFileReqDTO) error {

	sql :=
		`
			UPDATE files
			SET removed=true, client=$1
			WHERE file_name LIKE $2
				AND owner_id=(
							SELECT owner_id
							FROM owners
							WHERE folder=$3
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
func (p *postgresql) RestoreFile(ctx context.Context, restoreFileReqDTO daoDTO.RestoreFileReqDTO) error {

	sql :=
		`
		UPDATE files
		SET removed=false, client=$1
		WHERE file_name LIKE $2
			AND owner_id=(
						SELECT owner_id
						FROM owners
						WHERE folder=$3
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
func (p *postgresql) FindAllFilesByOwnerOrFileId(ctx context.Context, FindAllFilesByOwnerOrFileIdReqDTO daoDTO.FindAllFilesByOwnerOrFileIdReqDTO) (daoDTO.FindAllFilesByOwnerOrFileIdRespDTO, error) {

	sql :=
		`
		SELECT file_id, files.removed, virtual_name, state, hash_sum, file_name, size_file, files.owner_id, mod_time, client, username, folder
		FROM owners
		INNER JOIN files ON files.owner_id=owners.owner_id
		`

	var err error
	var rows pgx.Rows

	if FindAllFilesByOwnerOrFileIdReqDTO.OwnerId != 0 {
		sql = sql + " WHERE owners.owner_id=$1"
		rows, err = p.conn.Query(ctx, sql, FindAllFilesByOwnerOrFileIdReqDTO.OwnerId)
	} else if FindAllFilesByOwnerOrFileIdReqDTO.Folder != "" && FindAllFilesByOwnerOrFileIdReqDTO.Username != "" {
		sql = sql + " WHERE owners.folder=$1 AND owners.username=$2"
		rows, err = p.conn.Query(ctx, sql, FindAllFilesByOwnerOrFileIdReqDTO.Folder, FindAllFilesByOwnerOrFileIdReqDTO.Username)
	} else if FindAllFilesByOwnerOrFileIdReqDTO.FileId != 0 {
		sql = sql + " WHERE files.file_id=$1"
		rows, err = p.conn.Query(ctx, sql, FindAllFilesByOwnerOrFileIdReqDTO.FileId)
	} else {
		err = ErrInvalidOwner
	}
	if err != nil {
		return daoDTO.FindAllFilesByOwnerOrFileIdRespDTO{}, err
	}
	defer rows.Close()

	var files = make([]daoDTO.FullFile, 0)
	for rows.Next() {
		var file daoDTO.FullFile = daoDTO.FullFile{}
		err := rows.Scan(
			&file.FileId,
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
			return daoDTO.FindAllFilesByOwnerOrFileIdRespDTO{}, err
		}
		files = append(files, file)
	}
	return daoDTO.FindAllFilesByOwnerOrFileIdRespDTO{
		Files: files,
	}, nil
}

var _ usecase.IDaoFile = (*postgresql)(nil)
// var _ servic
