package _fileworker

import (
	"errors"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

var ErrNotFoundRealFolder = errors.New("ErrNotFoundRealFolder")

type fileWorker struct {
	mainDir string
}

func New(mainDir string) *fileWorker {

	return &fileWorker{
		mainDir: mainDir,
	}
}

func (f *fileWorker) writeBatch(path string, content *[]byte, batchSize int) error {

	file, err := os.Create(path)
	if err != nil {
		return err
	}

	n, err := file.Write(*content)
	if err != nil {
		return err
	}
	if n != batchSize {
		return errors.New("n != batchSize")
	}
	return nil
}

func (f *fileWorker) getFileName(path string) (string, error) {

	splitted := strings.Split(path, string(filepath.Separator))
	if len(splitted) == 0 {return "", errors.New("fileName is not exists")}
	return splitted[len(splitted)-1], nil
}

// TODO(дописать все аргументы)
func (f *fileWorker) getFullNameFile(version int, batch *Batch) string {

	fullName := strings.Join([]string{
		strconv.Itoa(version), batch.ModTime, batch.Hash,
		strconv.Itoa(batch.FullSize), strconv.Itoa(batch.Offset),
		strconv.Itoa(batch.Part), strconv.Itoa(batch.CountParts),
	}, "_")
	return fullName
}

func (f *fileWorker) getVersionFile(path string) (int, error) {

	files, err := f.getFilesInFolder(path)
	if err != nil {
		return -1, err
	}

	lastVersionFile := files[len(files)-1]
	versionStr := strings.Split(lastVersionFile, "_")[0]
	version, err := strconv.Atoi(versionStr)
	if err != nil {
		return -1, err
	}
	return version, nil
}

func (f *fileWorker) isFirstBatch(batch *Batch) bool {

	if batch.Part == 1 {return true} else {return false}
}

func (f *fileWorker) isLastBatch(batch *Batch) bool {

	if batch.Part == batch.CountParts {return true} else {return false}
}

func (f *fileWorker) createFolder(path string) error {

	if f.checkExistsFolder(path) == nil {
		return nil
	} else {
		return f.mkdirAll(path)
	}
}

func (f *fileWorker) checkExistsFolder(path string) error {
	
	_, err := os.Stat(path)
	return err
}

func (f *fileWorker) mkdirAll(path string) error {

	return os.MkdirAll(path, 0777)
}

func (f *fileWorker) getRealPath(startPath, path string) (string, error) {

	log.Println("path:", path)
	log.Println("startPath:", startPath)

	folders := strings.Split(path, string(filepath.Separator))

	log.Println("folders:", folders)

	realPath := startPath
	for _, folder := range folders {
		fileNames, err := f.getFilesInFolder(realPath)
		if err != nil {
			return realPath, err
		}

		log.Println("fileNames:", fileNames)

		fileName, ok := f.contains(fileNames, folder)
		if ok {
			realPath = filepath.Join(realPath, fileName)
		} else {
			return "", errors.New("Folder Is Not Exists")
		}
	}

	log.Println("realPath:", realPath)

	return realPath, nil
}

func (f *fileWorker) contains(m []string, e string) (string, bool) {

	for _, v := range m {
		if strings.Contains(v, e) {
			return v, true
		}
	}
	return "", false
}

func (f *fileWorker) checkFolder(folder string) bool {

	log.Println("folder:", folder)
	_, err := os.Stat(folder)
	if err == nil {
		return false
	}
	return false
}

func (f *fileWorker) getNameOfLastVersion(dir string) (string, error) {

	names, err := f.getFilesInFolder(dir)
	if err != nil {
		return "", err
	}

	if len(names) == 0 {
		return "", err
	}

	lastVersion := names[len(names)-1]

	log.Println("lastVersion:", lastVersion)

	return lastVersion, nil
}

func (f *fileWorker) getFilesInFolder(folder string) ([]string, error) {

	log.Println("folder:", folder)

	files, err := ioutil.ReadDir(folder)
	if err != nil {
		return nil, err
	}

	if len(files) == 0 {
		return nil, nil
	}

	log.Println("files:", files)

	names := []string{}
	for _, f := range files {
		names = append(names, f.Name())
	}

	sort.Strings(names)
	return names, nil
}
