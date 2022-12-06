package storage

import (
	"errors"
	"os"
	"path/filepath"

	repo "github.com/gobox-preegnees/file_manager/pkg/repo"
)

type storage struct{
	path string
}

func New() *storage {

	return &storage{}
}

var _ repo.IStorage = (*storage)(nil)

func (s *storage) SaveBatchOnDisk(path string, data *[]byte) (int, error) {

	path = filepath.Join(s.path, path)

	file, err := os.Create(path)
	if err != nil {
		os.Remove(path)
		return -1, err
	}
	defer file.Close()
	
	n, err := file.Write(*data)
	if err != nil {
		os.Remove(path)
		return -1, err
	}
	return n, nil
}

func (s *storage) GetBatchFromDisk(path string) (*[]byte, error) {

	path = filepath.Join(s.path, path)

	clear := func(err error) (*[]byte, error) {
		
		os.Remove(path)
		return nil, err
	}

	file, err := os.Open(path)
	if err != nil {
		return clear(err)
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		return clear(err)
	}

	data := make([]byte, stat.Size())
	n, err := file.Read(data)
	if err != nil {
		return clear(err)
	}
	if n == 0 {
		return nil, errors.New("n is 0")
	}
	return &data, nil
} 
