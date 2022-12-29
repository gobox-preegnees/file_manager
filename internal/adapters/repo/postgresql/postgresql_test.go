package postgresql

import (
	"context"
	"testing"

	repoDTO "github.com/gobox-preegnees/file_manager/internal/adapters/repo"
	"github.com/gobox-preegnees/file_manager/internal/domain/entity"
)

func TestFindAll(t *testing.T) {
	ctx := context.Background()
	url := "postgres://postgres:postgres@localhost:5431/postgres?sslmode=disable"
	connection, err := New(ctx, url)
	if err != nil {
		t.Fatal(err)
	}
	username:= "Albert_Robinson"
	folder := "folder1"
	files, err := connection.FindAll(ctx, username, folder)
	if err != nil {
		t.Error(err)
	}
	if len(files) == 0 {
		t.Error("len(files) == 0")
	} else {
		t.Log(files)
	}
}

func TestFindOneByPath(t *testing.T) {
	ctx := context.Background()
	url := "postgres://postgres:postgres@localhost:5431/postgres?sslmode=disable"
	connection, err := New(ctx, url)
	if err != nil {
		t.Error(err)
	}
	username:= "Albert_Robinson"
	folder := "folder1"
	fileName := "Photos/family.jpg"
	file, err := connection.FindOneByPath(ctx, username, folder, fileName)
	if err != nil {
		t.Error(err)
	}
	t.Log(file)
}

func TestSaveOne(t *testing.T) {
	ctx := context.Background()
	url := "postgres://postgres:postgres@localhost:5431/postgres?sslmode=disable"
	connection, err := New(ctx, url)
	if err != nil {
		t.Error(err)
	}
	username:= "Albert_Robinson"
	folder := "folder1"
	file := repoDTO.FileDTO{
		Client: "Win1",
		File: entity.File{
			FileName: "Photos/family.jpg",
			HashSum: "efrtg646567767868j",
			SizeFile: 1234,
			ModTime: 1646857394,
		},
	}
	err = connection.SaveOne(ctx, username, folder, file)
	if err != nil {
		t.Error(err)
	}
}

func TestUpdateState(t *testing.T) {
	ctx := context.Background()
	url := "postgres://postgres:postgres@localhost:5431/postgres?sslmode=disable"
	connection, err := New(ctx, url)
	if err != nil {
		t.Fatal(err)
	}
	username:= "Albert_Robinson"
	folder := "folder1"
	fileName := "Photos/family.jpg"
	virtualName := "newvirtNameLOL"
	state := 200
	file, err := connection.FindOneByPath(ctx, username, folder, fileName)
	if err != nil {
		t.Error(err)
	}
	t.Log(file)
	if file.State == 200 {
		state = 300
	}
	// TODO(обновление не происходит)
	err = connection.UpdateState(ctx, username, folder, file.Client, file.FileName, file.HashSum, virtualName, state)
	if err != nil {
		t.Error(err)
	}
}

func TestUpdateFileNameOneFile(t *testing.T) {
	ctx := context.Background()
	url := "postgres://postgres:postgres@localhost:5431/postgres?sslmode=disable"
	connection, err := New(ctx, url)
	if err != nil {
		t.Fatal(err)
	}
	username:= "Albert_Robinson"
	folder := "folder1"
	fileName := "Photos/family.jpg"
	file, err := connection.FindOneByPath(ctx, username, folder, fileName)
	if err != nil {
		t.Error(err)
	}
	t.Log(file)
	err = connection.UpdateFileName(ctx, username, folder, file.Client, file.FileName, file.FileName+"newName", file.HashSum)
	if err != nil {
		t.Error(err)
	}
}

func TestUpdateFileNameAllFiles(t *testing.T) {
	ctx := context.Background()
	url := "postgres://postgres:postgres@localhost:5431/postgres?sslmode=disable"
	connection, err := New(ctx, url)
	if err != nil {
		t.Fatal(err)
	}
	username:= "Albert_Robinson"
	folder := "folder1"
	fileName := "Photos"
	client := "cli1"
	err = connection.UpdateFileName(ctx, username, folder,client, fileName, "newName", "")
	if err != nil {
		t.Error(err)
	}
}