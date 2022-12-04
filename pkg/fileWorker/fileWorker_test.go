package fileworker

import (
	"testing"
	"os"
	"path/filepath"
)

const TEST_DIR = "TEST_DIR"

func TestGetPathLastVersion(t *testing.T) {

	os.MkdirAll(TEST_DIR, 0777) 
	defer os.RemoveAll(TEST_DIR)

	lastName := "3_name"
	folderNames := []string{"1_name", "2_name", lastName}
	for _, v := range folderNames {
		path := filepath.Join(TEST_DIR, v)
		os.Mkdir(path, 0777)
	}
	
	fileWorker := New(TEST_DIR)
	folderName, err := fileWorker.getNameOfLastVersion(TEST_DIR)
	if err != nil {
		t.Fatal(err)
	}
	if folderName != lastName {
		t.Fatal("folderName != lastName")
	}
}

func TestGetPathLastVersionIfFolderEmpty(t *testing.T) {
	
	os.MkdirAll(TEST_DIR, 0777) 
	defer os.RemoveAll(TEST_DIR)
	
	fileWorker := New(TEST_DIR)
	folderName, err := fileWorker.getNameOfLastVersion(TEST_DIR)
	if err != nil {
		t.Fatal(err)
	}
	if folderName != "" {
		t.Fatal("folderName != \"\"")
	}
}

func TestGetFullFolderPath(t *testing.T) {

	os.MkdirAll(TEST_DIR, 0777) 
	defer os.RemoveAll(TEST_DIR)

	folder1 := "Folder1_12345"
	folder2 := "Folder2_23456"
	folder3 := "Folder3_98765"
	os.MkdirAll(filepath.Join(TEST_DIR, folder1), 0777)
	os.MkdirAll(filepath.Join(TEST_DIR, folder2), 0777)
	os.MkdirAll(filepath.Join(TEST_DIR, folder2, folder3), 0777)

	folderPath := filepath.Join("Folder2", "Folder3")
	realPath := filepath.Join(TEST_DIR, folder2, folder3)

	fileWorker := New(TEST_DIR)
	path, err := fileWorker.getRealPath(TEST_DIR, folderPath)
	if err != nil {
		t.Fatal(err)
	}
	if realPath != path {
		t.Fatal("realPath != path")
	}
}

func TestGetFullFolderPathIfNotExistsFolder(t *testing.T) {

	os.MkdirAll(TEST_DIR, 0777) 
	defer os.RemoveAll(TEST_DIR)

	folder1 := "Folder1_12345"
	folder2 := "Folder2_23456"
	folder3 := "Folder3_46576"
	os.MkdirAll(filepath.Join(TEST_DIR, folder1), 0777)
	os.MkdirAll(filepath.Join(TEST_DIR, folder2), 0777)

	folderPath := filepath.Join("Folder2", "Folder3")
	path := filepath.Join(TEST_DIR, folder2, folder3)

	fileWorker := New(TEST_DIR)
	realPath, err := fileWorker.getRealPath(TEST_DIR, folderPath)
	if err == nil {
		t.Fatal(err)
	}
	if realPath == path {
		t.Fatal("realPath != path")
	}
}
