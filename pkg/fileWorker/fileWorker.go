package fileworker

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sort"
)

type fileWorker struct {
	path string
}

func New() *fileWorker {

	return &fileWorker{}
}

func (f *fileWorker) Save(batch BatchForSave) error {

	// filepath.Join(f.path, )
	return nil
}

func (f *fileWorker) getFullPath(parent, path string) string {

	fullPath := filepath.Join(parent, path)
	dir := filepath.Dir(fullPath)
	if !f.checkFolder(dir) {

	}
	return ""
}

func (f *fileWorker) checkFolder(dir string) bool {
	
	_, err := os.Stat(dir)
	if err == nil {
		return false
	}
	return false
}

func (f *fileWorker) getNameLastVersion(dir string) (string, error) {

	log.Println("dir:", dir)

	files, err := ioutil.ReadDir(dir)
    if err != nil {
        return "", err
    }

	if len(files) == 0 {
		return "", nil
	}

	log.Println("files:", files)
 
	names := []string{}
    for _, f := range files {
        names = append(names, f.Name())
    }

	sort.Strings(names)
	lastVersion := names[len(names)-1]

	log.Println("lastVersion:", lastVersion)
	return lastVersion, nil
}