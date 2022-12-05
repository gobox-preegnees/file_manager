package repo

import (
	"context"
	"testing"
)

func TestConnToDB(t *testing.T) {

	url := "user=docker password=docker host=localhost port=5431 dbname=docker"
	_, err := New(nil, url)
	if err != nil {
		t.Fatal(err)
	}
}

func TestSaveBatch(t *testing.T) {

	url := "user=docker password=docker host=localhost port=5431 dbname=docker"
	repo, err := New(nil, url)
	if err != nil {
		t.Fatal(err)
	}
	batch := Batch{
		Username: "usrename",
		FolderID: "folder",
		ClientID: "client",
		Path: "/folder/file.txt",
		Hash: "frekth46",
		ModTime: 12347,
		Part: 2,
		CountParts: 10,
		PartSize: 2048,
		Offset: 4096,
		SizeFile: 100000,
		Content: []byte("hello world"),
	} 
	id, err := repo.SaveBatch(context.TODO(), &batch)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(id)
}