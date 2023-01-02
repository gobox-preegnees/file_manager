package postgresql

import (
	"context"
	"os"
	"testing"

	daoDTO "github.com/gobox-preegnees/file_manager/internal/adapters/dao"
)

var postgres *postgresql
var ctx context.Context = context.Background()

const url string = "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"

func TestMain(m *testing.M) {

	var err error
	postgres, err = NewPosgresql(ctx, url)
	if err != nil {
		panic(err)
	}
	code := m.Run()
	postgres.conn.Close(ctx)
	os.Exit(code)
}

func truncate() {
	_, err := postgres.conn.Exec(ctx, "TRUNCATE TABLE owners CASCADE")
	if err != nil {
		panic(err)
	}
}

func TestSaveOwner(t *testing.T) {

	truncate()
	defer truncate()

	data := []struct {
		Username string
		Folder   string
		Err      bool
	}{
		{
			Username: "u1",
			Folder:   "f1",
			Err:      false,
		},
		{
			Username: "u1",
			Folder:   "f1",
			Err:      true,
		},
		{
			Username: "u2",
			Folder:   "f2",
			Err:      false,
		},
	}

	for _, d := range data {
		t.Run("test", func(t *testing.T) {
			id, err := postgres.SaveOwner(ctx, daoDTO.SaveOwnerDTO{
				Identifier: daoDTO.Identifier{
					Username: d.Username,
					Folder:   d.Folder,
				},
			})
			if d.Err {
				if err == nil {
					t.Errorf("expected error {%v}, got none", err)
				}
			}
			if id == 0 && !d.Err {
				t.Errorf("expected id, got none")
			}
			t.Logf("id=%d", id)
		})
	}
}

func TestRenameOwner(t *testing.T) {

	truncate()
	defer truncate()

	data := []struct {
		// Create
		Username string
		Folder   string

		// Rename
		NewName string
		Err     bool
	}{
		{
			Username: "u1",
			Folder:   "f1",

			NewName: "f2",
			Err:     false,
		},
	}

	for _, d := range data {
		t.Run("test", func(t *testing.T) {
			id, err := postgres.SaveOwner(ctx, daoDTO.SaveOwnerDTO{
				Identifier: daoDTO.Identifier{
					Username: d.Username,
					Folder:   d.Folder,
				},
			})
			if err != nil {
				t.Errorf("expected none, got error {%v} in create", err)
			}
			if id == 0 {
				t.Errorf("expected id, got none")
			}
			t.Logf("id=%d", id)

			err = postgres.RenameOwner(ctx, daoDTO.RenameOwnerDTO{
				OwnerID: id,
				NewName: d.NewName,
			})
			if err != nil {
				if !d.Err {
					t.Errorf("expected none, got error {%v}", err)
				}
			}
		})
	}
}

func TestDeleteOwner(t *testing.T) {

	truncate()
	defer truncate()

	data := []struct {
		// Create
		Username string
		Folder   string

		// Delete
		Err bool
	}{
		{
			Username: "u1",
			Folder:   "f1",

			Err: false,
		},
	}

	for _, d := range data {
		t.Run("test", func(t *testing.T) {
			id, err := postgres.SaveOwner(ctx, daoDTO.SaveOwnerDTO{
				Identifier: daoDTO.Identifier{
					Username: d.Username,
					Folder:   d.Folder,
				},
			})
			if err != nil {
				t.Errorf("expected none, got error {%v} in create", err)
			}
			if id == 0 {
				t.Errorf("expected id, got none")
			}
			t.Logf("id=%d", id)

			err = postgres.DeleteOwner(ctx, daoDTO.DeleteOwnerDTO{
				OwnerID: id,
			})
			if err != nil {
				if !d.Err {
					t.Errorf("expected none, got error {%v}", err)
				}
			}
		})
	}
}

func TestFindAllOwners(t *testing.T) {

	truncate()
	defer truncate()

	data := []struct {
		// Create
		Username string
		Folder   string

		// Find
		Count int
		Err   bool
	}{
		{
			Username: "u1",
			Folder:   "f1",

			Count: 1,
			Err:   false,
		},
		{
			Username: "u1",
			Folder:   "f2",

			Count: 2,
			Err:   false,
		},
		{
			Username: "u2",
			Folder:   "f1",

			Count: 1,
			Err:   false,
		},
	}

	for _, d := range data {
		t.Run("test", func(t *testing.T) {
			id, err := postgres.SaveOwner(ctx, daoDTO.SaveOwnerDTO{
				Identifier: daoDTO.Identifier{
					Username: d.Username,
					Folder:   d.Folder,
				},
			})
			if err != nil {
				t.Errorf("expected none, got error {%v} in create", err)
			}
			if id == 0 {
				t.Errorf("expected id, got none")
			}
			t.Logf("id=%d", id)

			owners, err := postgres.FindAllOwners(ctx, daoDTO.FindAllOwnersReqDTO{
				Username: d.Username,
			})
			if err != nil {
				if !d.Err {
					t.Errorf("expected none, got error {%v}", err)
				}
			}
			if len(owners.Owners) != d.Count {
				t.Errorf("expected %d, got %d, username %s", d.Count, len(owners.Owners), d.Username)
			}
		})
	}
}

func TestSaveFile(t *testing.T) {

	truncate()
	defer truncate()

	data := []struct {
		daoDTO.Identifier
		daoDTO.File
		ClientId string
		Err      bool
	}{
		{
			Identifier: daoDTO.Identifier{
				Username: "1",
				Folder:   "1",
			},
			File: daoDTO.File{
				FileName: "/tmp/test.txt",
				SizeFile: 1,
				HashSum:  "1",
				ModTime:  1,
			},
			ClientId: "1",
			Err:      false,
		},
		{
			Identifier: daoDTO.Identifier{
				Username: "1",
				Folder:   "1",
			},
			File: daoDTO.File{
				FileName: "/tmp/test.txt",
				SizeFile: 1,
				HashSum:  "1",
				ModTime:  1,
			},
			ClientId: "1",
			Err:      true,
		},
	}

	for _, d := range data {
		t.Run("test", func(t *testing.T) {
			_, _ = postgres.SaveOwner(ctx, daoDTO.SaveOwnerDTO{
				Identifier: d.Identifier,
			})

			id, err := postgres.SaveFile(ctx, daoDTO.SaveFileReqDTO{
				Identifier: d.Identifier,
				File:       d.File,
				Client:     d.ClientId,
			})
			if err != nil {
				if !d.Err {
					t.Errorf("expected none, got error {%v} in create", err)
				}
			}
			if id == 0 {
				if !d.Err {
					t.Errorf("expected id, got none")
				}
			}
			t.Logf("%d", id)
		})
	}
}

func TestSetState(t *testing.T) {

	truncate()
	defer truncate()

	data := []struct {
		daoDTO.Identifier
		daoDTO.File
		ClientId    string
		VirtualName string
		State       int
		Err         bool
	}{
		{
			Identifier: daoDTO.Identifier{
				Username: "1",
				Folder:   "1",
			},
			File: daoDTO.File{
				FileName: "/tmp/test.txt",
				SizeFile: 1,
				HashSum:  "1",
				ModTime:  1,
			},
			ClientId:    "1",
			VirtualName: "1",
			State:       300,
			Err:         false,
		},
		{
			Identifier: daoDTO.Identifier{
				Username: "1",
				Folder:   "1",
			},
			File: daoDTO.File{
				FileName: "/tmp/test.txt",
				SizeFile: 1,
				HashSum:  "1",
				ModTime:  1,
			},
			ClientId:    "1",
			VirtualName: "1",
			State:       100,
			Err:         false,
		},
		{
			Identifier: daoDTO.Identifier{
				Username: "1",
				Folder:   "1",
			},
			File: daoDTO.File{
				FileName: "invalid name",
				SizeFile: 1,
				HashSum:  "1",
				ModTime:  1,
			},
			ClientId:    "1",
			VirtualName: "1",
			State:       200,
			Err:         true,
		},
	}

	for _, d := range data {
		t.Run("test", func(t *testing.T) {
			_, _ = postgres.SaveOwner(ctx, daoDTO.SaveOwnerDTO{
				Identifier: d.Identifier,
			})

			_, _ = postgres.SaveFile(ctx, daoDTO.SaveFileReqDTO{
				Identifier: d.Identifier,
				File:       d.File,
				Client:     d.ClientId,
			})

			err := postgres.SetState(ctx, daoDTO.SetStateReqDTO{
				Identifier: daoDTO.Identifier{
					Username: d.Username,
					Folder:   d.Folder,
				},
				FileName:    d.File.FileName,
				VirtualName: d.VirtualName,
				State:       d.State,
			})

			if err != nil {
				if !d.Err {
					t.Errorf("expected none, got error {%v} in create", err)
				}
			}
		})
	}
}

func TestRenameFile(t *testing.T) {

	truncate()
	defer truncate()

	data := []struct {
		daoDTO.Identifier
		files    []daoDTO.File
		ClientId string
		NewName  string
		OldName  string
		Err      bool
	}{
		{
			Identifier: daoDTO.Identifier{
				Username: "1",
				Folder:   "1",
			},
			files: []daoDTO.File{
				{
					FileName: "folder/test.txt",
					SizeFile: 1,
					HashSum:  "1",
					ModTime:  1,
				},
			},
			ClientId: "1",
			NewName:  "FirstTest/newName.txt",
			OldName:  "folder/test.txt",
			Err:      false,
		},
		{
			Identifier: daoDTO.Identifier{
				Username: "1",
				Folder:   "1",
			},
			files: []daoDTO.File{
				{
					FileName: "not/test.txt",
					SizeFile: 2,
					HashSum:  "2",
					ModTime:  2,
				},
				{
					FileName: "tmp/test1.abd",
					SizeFile: 3,
					HashSum:  "3",
					ModTime:  3,
				},
				{
					FileName: "tmp/var/.txt",
					SizeFile: 4,
					HashSum:  "4",
					ModTime:  4,
				},
				{
					FileName: "tmp/",
					SizeFile: 0,
					HashSum:  "",
					ModTime:  4,
				},
			},
			ClientId: "2",
			NewName:  "newName/",
			OldName:  "tmp/",
			Err:      false,
		},
		{
			Identifier: daoDTO.Identifier{
				Username: "1",
				Folder:   "1",
			},
			files: []daoDTO.File{
				{
					FileName: "11not/test.txt",
					SizeFile: 3,
					HashSum:  "3",
					ModTime:  3,
				},
				{
					FileName: "tmp/folder/test1.abd",
					SizeFile: 3,
					HashSum:  "3",
					ModTime:  3,
				},
				{
					FileName: "tmp/folder/var/.txt",
					SizeFile: 4,
					HashSum:  "4",
					ModTime:  4,
				},
			},
			ClientId: "2",
			NewName:  "newName/LOLOLOL/",
			OldName:  "tmp/folder/",
			Err:      false,
		},
	}

	for _, d := range data {
		t.Run("test", func(t *testing.T) {
			_, _ = postgres.SaveOwner(ctx, daoDTO.SaveOwnerDTO{
				Identifier: d.Identifier,
			})

			for _, f := range d.files {
				_, err := postgres.SaveFile(ctx, daoDTO.SaveFileReqDTO{
					Identifier: d.Identifier,
					File:       f,
					Client:     d.ClientId,
				})
				if err != nil {
					t.Fatalf("Failed to save, error {%v}", err)
				}
			}

			err := postgres.RenameFile(ctx, daoDTO.RenameFileReqDTO{
				Identifier: daoDTO.Identifier{
					Username: d.Username,
					Folder:   d.Folder,
				},
				Client:  d.ClientId,
				OldName: d.OldName,
				NewName: d.NewName,
			})

			if err != nil {
				if !d.Err {
					t.Errorf("expected none, got error {%v} in create", err)
				}
			}
		})
	}
}

func TestDeleteFileAndRestoreFile(t *testing.T) {

	truncate()
	defer truncate()

	data := []struct {
		daoDTO.Identifier
		files    []daoDTO.File
		ClientId string
		FileName string
		Err      bool
	}{
		{
			Identifier: daoDTO.Identifier{
				Username: "username",
				Folder:   "1",
			},
			files: []daoDTO.File{
				{
					FileName: "folder/test.txt",
					SizeFile: 1,
					HashSum:  "1",
					ModTime:  1,
				},
			},
			ClientId: "1",
			FileName: "folder/test.txt",
			Err:      false,
		},
		{
			Identifier: daoDTO.Identifier{
				Username: "username",
				Folder:   "1",
			},
			files: []daoDTO.File{
				{
					FileName: "not/test.txt",
					SizeFile: 2,
					HashSum:  "2",
					ModTime:  2,
				},
				{
					FileName: "tmp/test1.abd",
					SizeFile: 3,
					HashSum:  "3",
					ModTime:  3,
				},
				{
					FileName: "tmp/var/.txt",
					SizeFile: 4,
					HashSum:  "4",
					ModTime:  4,
				},
				{
					FileName: "tmp/",
					SizeFile: 0,
					HashSum:  "",
					ModTime:  4,
				},
			},
			ClientId: "2",
			FileName: "tmp/",
			Err:      false,
		},
		{
			Identifier: daoDTO.Identifier{
				Username: "username",
				Folder:   "1",
			},
			files: []daoDTO.File{
				{
					FileName: "11not/test.txt",
					SizeFile: 3,
					HashSum:  "3",
					ModTime:  3,
				},
				{
					FileName: "tmp/folder/test1.abd",
					SizeFile: 3,
					HashSum:  "3",
					ModTime:  3,
				},
				{
					FileName: "tmp/folder/var/.txt",
					SizeFile: 4,
					HashSum:  "4",
					ModTime:  4,
				},
			},
			ClientId: "2",
			FileName: "tmp/folder/",
			Err:      false,
		},
	}

	for _, d := range data {
		t.Run("test", func(t *testing.T) {
			_, _ = postgres.SaveOwner(ctx, daoDTO.SaveOwnerDTO{
				Identifier: d.Identifier,
			})

			for _, f := range d.files {
				_, err := postgres.SaveFile(ctx, daoDTO.SaveFileReqDTO{
					Identifier: d.Identifier,
					File:       f,
					Client:     d.ClientId,
				})
				if err != nil {
					t.Fatalf("Failed to save file, error {%v}", err)
				}
			}

			err := postgres.DeleteFile(ctx, daoDTO.DeleteFileReqDTO{
				Identifier: daoDTO.Identifier{
					Username: d.Username,
					Folder:   d.Folder,
				},
				Client:   d.ClientId,
				FileName: d.FileName,
			})

			if err != nil {
				if !d.Err {
					t.Errorf("expected none, got error {%v} in create", err)
				}
			}

			err = postgres.RestoreFile(ctx, daoDTO.RestoreFileReqDTO{
				Identifier: daoDTO.Identifier{
					Username: d.Username,
					Folder:   d.Folder,
				},
				Client:   d.ClientId,
				FileName: d.FileName,
			})

			if err != nil {
				if !d.Err {
					t.Errorf("expected none, got error {%v} in create", err)
				}
			}
		})
	}
}

func TestFindAllFilesByOwner(t *testing.T) {

	truncate()
	defer truncate()

	data := []struct {
		Identifier          daoDTO.Identifier
		files               []daoDTO.File
		Client              string
		ByOwnerId           bool
		ByFileId            bool
		ByUsernameAndFolder bool
		Err                 bool
		Count               int
	}{
		{
			files: []daoDTO.File{
				{
					FileName: "11not/test.txt",
					SizeFile: 3,
					HashSum:  "3",
					ModTime:  3,
				},
				{
					FileName: "tmp/folder/test1.abd",
					SizeFile: 4,
					HashSum:  "4",
					ModTime:  4,
				},
				{
					FileName: "tmp/folder/var/.txt",
					SizeFile: 5,
					HashSum:  "5",
					ModTime:  5,
				},
			},
			Identifier: daoDTO.Identifier{
				Username: "username",
				Folder:   "folder",
			},
			Client:              "1",
			Count:               3,
			ByOwnerId:           false,
			ByFileId:            false,
			ByUsernameAndFolder: true,
			Err:                 false,
		},
		{
			files: []daoDTO.File{
				{
					FileName: "vyub67u67/test.txt",
					SizeFile: 30,
					HashSum:  "30",
					ModTime:  30,
				},
				{
					FileName: "tmp/7u67/test1.abd",
					SizeFile: 40,
					HashSum:  "40",
					ModTime:  40,
				},
				{
					FileName: "tmp/folder/o9p0/.txt",
					SizeFile: 50,
					HashSum:  "50",
					ModTime:  50,
				},
			},
			Identifier: daoDTO.Identifier{
				Username: "username",
				Folder:   "folder",
			},
			Client:              "1",
			Count:               6,
			ByOwnerId:           true,
			ByFileId:            false,
			ByUsernameAndFolder: false,
			Err:                 false,
		},
		{
			files: []daoDTO.File{
				{
					FileName: "abcdef",
					SizeFile: 6,
					HashSum:  "6",
					ModTime:  6,
				},
				{
					FileName: "adergrbu",
					SizeFile: 7,
					HashSum:  "7",
					ModTime:  7,
				},
			},
			Identifier: daoDTO.Identifier{
				Username: "username",
				Folder:   "folder",
			},
			Client:              "2",
			Count:               1,
			ByOwnerId:           false,
			ByFileId:            true,
			ByUsernameAndFolder: false,
			Err:                 false,
		},
	}

	ownerID, err := postgres.SaveOwner(ctx, daoDTO.SaveOwnerDTO{
		Identifier: data[0].Identifier,
	})
	if err != nil {
		t.Fatalf("expected no error, got error {%v} in save owner", err)
	}

	for _, d := range data {
		t.Run("test", func(t *testing.T) {
			var fileIDs = make([]int, 0)
			for _, f := range d.files {
				fileID, err := postgres.SaveFile(ctx, daoDTO.SaveFileReqDTO{
					Identifier: d.Identifier,
					File:       f,
					Client:     d.Client,
				})
				if err != nil {
					t.Fatalf("Failed to save, error {%v}", err)
				}
				fileIDs = append(fileIDs, fileID)
			}

			var err error
			var files daoDTO.FindAllFilesByOwnerOrFileIdRespDTO

			if d.ByOwnerId {
				files, err = postgres.FindAllFilesByOwnerOrFileId(ctx, daoDTO.FindAllFilesByOwnerOrFileIdReqDTO{
					Owner: daoDTO.Owner{
						OwnerId: ownerID,
					},
				})
			} else if d.ByUsernameAndFolder {
				files, err = postgres.FindAllFilesByOwnerOrFileId(ctx, daoDTO.FindAllFilesByOwnerOrFileIdReqDTO{
					Owner: daoDTO.Owner{
						Identifier: d.Identifier,
					},
				})
			} else if d.ByFileId {
				files, err = postgres.FindAllFilesByOwnerOrFileId(ctx, daoDTO.FindAllFilesByOwnerOrFileIdReqDTO{
					FileId: fileIDs[0],
				})
			}

			if err != nil {
				if !d.Err {
					t.Errorf("expected none, got error {%v} in findAllFilesByOwner", err)
				}
			}
			if len(files.Files) != d.Count {
				t.Errorf("expected count {%v}, got {%v}", d.Count, len(files.Files))
			}
		})
	}
}
