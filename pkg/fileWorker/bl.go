package fileworker

import (
	"path/filepath"
	"strconv"
)

func (f *fileWorker) SaveBatch(batch *Batch) error {

	startPath := filepath.Join(f.mainDir, batch.Username, batch.FolderID)

	if err := f.createFolder(startPath); err != nil {
		return err
	}

	realPath, err := f.getRealPath(startPath, filepath.Dir(batch.Path))
	if err != nil {
		return err
	}

	fileName, err := f.getFileName(batch.Path)
	if err != nil {
		return err
	}

	fullPath := filepath.Join(realPath, fileName)

	if !f.checkFolder(fullPath) {
		if err := f.createFolder(fullPath); err != nil {
			return err
		}
	}

	version := 1
	fullName := filepath.Join(fullPath, fileName)
	if f.isFirstBatch(batch) {
		version, err = f.getVersionFile(batch.Path)
		if err != nil {
			return err
		} else {
			version++
			fileName := f.getFullNameFile(version, batch)
			fullName = filepath.Join(fullPath, fileName, "_.")
			if err := f.createFolder(fullName); err != nil {
				return err
			}
		}
	}

	fileBatchName := filepath.Join(fullName, strconv.Itoa(batch.Part))
	if err := f.writeBatch(fileBatchName, &batch.Content, batch.PartSize); err != nil {
		return err
	}

	if f.isLastBatch(batch) {

	}

	return nil
}

// func (f *fileWorker) SaveBatch(batch Batch) error {

// 	startPath := f.join(f.mainDir, batch.Username, batch.FolderID)

// 	if err := f.createFolder(startPath); err != nil {
// 		return err
// 	}

// 	_, err := f.getRealPath(startPath, filepath.Dir(batch.Path))
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
