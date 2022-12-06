package repo

import (
	"context"
	"errors"
	"log"
	"strconv"

	"github.com/jackc/pgx/v5"
)

//go:generate mockgen -destination=../mocks/IStorage.go -source=repo.go
type IStorage interface {
	SaveBatchOnDisk(path string, data *[]byte) (n int, err error)
	GetBatchFromDisk(path string) (data *[]byte, err error)
}

type repo struct {
	conn    *pgx.Conn
	storage IStorage
}

func New(storage IStorage, url string) (*repo, error) {

	conn, err := pgx.Connect(context.Background(), url)
	if err != nil {
		return nil, err
	}
	if err := conn.Ping(context.TODO()); err != nil {
		return nil, err
	}

	return &repo{
		conn:    conn,
		storage: storage,
	}, nil
}

func (r *repo) SaveBatch(ctx context.Context, batch *Batch) (int, error) {

	sql := `INSERT INTO batches (
		username, folder, client, path, hash, mod_time, 
		part, count_parts, part_size, byte_offset, size_file 
	) 
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
	RETURNING batch_id;`

	tx, err := r.conn.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return -1, err
	}

	var batch_id int
	if err := tx.QueryRow(
		ctx, sql,
		batch.Username, batch.FolderID, batch.ClientID, batch.Path, batch.Hash, batch.ModTime,
		batch.Part, batch.CountParts, batch.PartSize, batch.Offset, batch.SizeFile,
	).Scan(&batch_id); err != nil {
		if err := tx.Rollback(ctx); err != nil {
			return -1, err
		}
		return -1, err
	}

	log.Println(batch_id)

	n, err := r.storage.SaveBatchOnDisk(strconv.Itoa(batch_id), &batch.Content)
	if n != batch.PartSize {
		return -1, errors.New("size batch is not equal saved batch")
	}

	if err != nil {
		if err := tx.Rollback(ctx); err != nil {
			return -1, err
		}
		return -1, err // TODO(что то сделать с n != batch.PartSize, сделать ошибку для этого)
	}

	if err := tx.Commit(ctx); err != nil {
		if err := tx.Rollback(ctx); err != nil {
			return -1, err
		}
		return -1, err
	}

	return batch_id, nil
}

func (r *repo) GetBatch(ctx context.Context, id int) (*Batch, error) {

	sql := `
	SELECT 
	username, folder, client, path, hash, 
	mod_time, part, count_parts, part_size, byte_offset, size_file
	FROM batches
	WHERE batch_id=$1;
	`
	batch := Batch{} 
	err := r.conn.QueryRow(ctx, sql, id).Scan(
		&batch.Username, &batch.FolderID, &batch.ClientID, &batch.Path, &batch.Hash, &batch.ModTime,
		&batch.Part, &batch.CountParts, &batch.PartSize, &batch.Offset, &batch.SizeFile,
	)
	if err != nil {
		return nil, err
	}
	// TODO(тут еще нужно получать файл)
	return &batch, nil
}

func (r *repo) deleteTestData(sql string) {

	r.conn.Exec(context.TODO(), sql)
}
