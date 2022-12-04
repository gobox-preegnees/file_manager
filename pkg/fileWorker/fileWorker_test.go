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
	
	fileWorker := New()
	folderName, err := fileWorker.getNameLastVersion(TEST_DIR)
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
	
	fileWorker := New()
	folderName, err := fileWorker.getNameLastVersion(TEST_DIR)
	if err != nil {
		t.Fatal(err)
	}
	if folderName != "" {
		t.Fatal("folderName != \"\"")
	}
}