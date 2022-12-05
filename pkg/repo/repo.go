package repo

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
)

type IStorage interface {
	SaveBatchOnDisk(path string, data *[]byte) (err error)
	GetBatchFromDisk(path string) (data *[]byte, err error)
}

type repo struct{
	conn *pgx.Conn
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
		conn: conn,
		storage: storage,
	}, nil
}

func (r *repo) SaveBatch(ctx context.Context, batch *Batch) (id int, err error) {

	sql := `INSERT INTO batches (
		username, folder, client, path, hash, mod_time, 
		part, count_parts, part_size, byte_offset, size_file 
	) 
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
	RETURNING batch_id;`

	var batch_id int
	err = r.conn.QueryRow(
		ctx, sql, 
		batch.Username, batch.FolderID, batch.ClientID, batch.Path, batch.Hash, batch.ModTime,
		batch.Part, batch.CountParts, batch.PartSize, batch.Offset, batch.SizeFile,
	).Scan(&batch_id)

	log.Println(batch_id)

	// r.conn.Exec(ctx, `TRUNCATE TABLE batches;`)

	if err != nil {
		return 0, err
	}

	return id, nil
}