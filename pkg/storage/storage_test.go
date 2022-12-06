package storage

import (
	"os"
	"testing"
)

const TEST_DIR = "TEST_DIR"
func TestSaveBat(t *testing.T) {

	os.MkdirAll(TEST_DIR, 0777)
	defer os.RemoveAll(TEST_DIR)
	
	s := New()
	s.path = TEST_DIR
	fileName := "1"
	data := []byte("hello world")
	_, err := s.SaveBatchOnDisk(fileName, &data)
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetBatch(t *testing.T) {

	os.MkdirAll(TEST_DIR, 0777)
	defer os.RemoveAll(TEST_DIR)
	
	s := New()
	s.path = TEST_DIR
	fileName := "1"
	data := []byte("hello world")
	s.SaveBatchOnDisk(fileName, &data)

	d, err := s.GetBatchFromDisk(fileName)
	if err != nil {
		t.Fatal(err)
	}
	if string(*d) != string(data) {
		t.Fatal("string(*d) != string(data)")
	}
}