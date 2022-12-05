package fileworker

import "context"

type IRepo interface {
	SaveBatch(ctx context.Context, batch *Batch) (id int, err error)
	GetBatch(ctx context.Context, id int) (batch *Batch, err error)
	Rename(ctx context.Context, oldPath, newPath string) (err error)
	Delete(ctx context.Context, path string) (err error)
	Recover(ctx context.Context, path string) (err error)
}

type fileWorker struct {
	repo IRepo
}

func (f *fileWorker) SaveBatch(ctx context.Context, batch *Batch) error {

	f.repo.SaveBatch(ctx, batch)
	return nil
}
