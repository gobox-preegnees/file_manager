package postgresql

import (
	"context"
	"os"
	// "strconv"
	repoDTO "github.com/gobox-preegnees/file_manager/internal/adapters/repo"
	"testing"
	// "github.com/gobox-preegnees/file_manager/internal/domain/entity"
)

var postgres *postgresql
var ctx context.Context = context.Background()

const url string = "postgres://postgres:postgres@localhost:5431/postgres?sslmode=disable"

func TestMain(m *testing.M) {

	var err error
	postgres, err = New(ctx, url)
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
			id, err := postgres.SaveOwner(ctx, repoDTO.SaveOwnerDTO{
				Identifier: repoDTO.Identifier{
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
			id, err := postgres.SaveOwner(ctx, repoDTO.SaveOwnerDTO{
				Identifier: repoDTO.Identifier{
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

			err = postgres.RenameOwner(ctx, repoDTO.RenameOwnerDTO{
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
			id, err := postgres.SaveOwner(ctx, repoDTO.SaveOwnerDTO{
				Identifier: repoDTO.Identifier{
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

			err = postgres.DeleteOwner(ctx, repoDTO.DeleteOwnerDTO{
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
			id, err := postgres.SaveOwner(ctx, repoDTO.SaveOwnerDTO{
				Identifier: repoDTO.Identifier{
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

			owners, err := postgres.FindAllOwners(ctx, repoDTO.FindAllOwnersReqDTO{
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
		repoDTO.Identifier
		repoDTO.File
		ClientId string
		Err      bool
	}{
		{
			Identifier: repoDTO.Identifier{
				Username: "1",
				Folder:   "1",
			},
			File: repoDTO.File{
				FileName: "/tmp/test.txt",
				SizeFile: 1,
				HashSum:  "1",
				ModTime:  1,
			},
			ClientId: "1",
			Err:      false,
		},
		{
			Identifier: repoDTO.Identifier{
				Username: "1",
				Folder:   "1",
			},
			File: repoDTO.File{
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
			_, _ = postgres.SaveOwner(ctx, repoDTO.SaveOwnerDTO{
				Identifier: d.Identifier,
			})

			id, err := postgres.SaveFile(ctx, repoDTO.SaveFileReqDTO{
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
		repoDTO.Identifier
		repoDTO.File
		ClientId    string
		VirtualName string
		State       int
		Err         bool
	}{
		{
			Identifier: repoDTO.Identifier{
				Username: "1",
				Folder:   "1",
			},
			File: repoDTO.File{
				FileName: "/tmp/test.txt",
				SizeFile: 1,
				HashSum:  "1",
				ModTime:  1,
			},
			ClientId:    "1",
			VirtualName: "1",
			State:       200,
			Err:         false,
		},
		{
			Identifier: repoDTO.Identifier{
				Username: "1",
				Folder:   "1",
			},
			File: repoDTO.File{
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
			_, _ = postgres.SaveOwner(ctx, repoDTO.SaveOwnerDTO{
				Identifier: d.Identifier,
			})

			_, _ = postgres.SaveFile(ctx, repoDTO.SaveFileReqDTO{
				Identifier: d.Identifier,
				File:       d.File,
				Client:     d.ClientId,
			})

			err := postgres.SetState(ctx, repoDTO.SetStateReqDTO{
				Identifier: repoDTO.Identifier{
					Username: d.Username,
					Folder:   d.Folder,
				},
				File:        d.File,
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
		repoDTO.Identifier
		files    []repoDTO.File
		ClientId string
		NewName  string
		OldName  string
		Err      bool
	}{
		{
			Identifier: repoDTO.Identifier{
				Username: "1",
				Folder:   "1",
			},
			files: []repoDTO.File{
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
			Identifier: repoDTO.Identifier{
				Username: "1",
				Folder:   "1",
			},
			files: []repoDTO.File{
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
			Identifier: repoDTO.Identifier{
				Username: "1",
				Folder:   "1",
			},
			files: []repoDTO.File{
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
			_, _ = postgres.SaveOwner(ctx, repoDTO.SaveOwnerDTO{
				Identifier: d.Identifier,
			})

			for _, f := range d.files {
				_, err := postgres.SaveFile(ctx, repoDTO.SaveFileReqDTO{
					Identifier: d.Identifier,
					File:       f,
					Client:     d.ClientId,
				})
				if err != nil {
				    t.Fatalf("Failed to save, error {%v}", err)
				}
			}

			err := postgres.RenameFile(ctx, repoDTO.RenameFileReqDTO{
				Identifier: repoDTO.Identifier{
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
		repoDTO.Identifier
		files    []repoDTO.File
		ClientId string
		FileName  string
		Err      bool
	}{
		{
			Identifier: repoDTO.Identifier{
				Username: "username",
				Folder:   "1",
			},
			files: []repoDTO.File{
				{
					FileName: "folder/test.txt",
					SizeFile: 1,
					HashSum:  "1",
					ModTime:  1,
				},
			},
			ClientId: "1",
			FileName:  "folder/test.txt",
			Err:      false,
		},
		{
			Identifier: repoDTO.Identifier{
				Username: "username",
				Folder:   "1",
			},
			files: []repoDTO.File{
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
			FileName:  "tmp/",
			Err:      false,
		},
		{
			Identifier: repoDTO.Identifier{
				Username: "username",
				Folder:   "1",
			},
			files: []repoDTO.File{
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
			FileName:  "tmp/folder/",
			Err:      false,
		},
	}

	for _, d := range data {
		t.Run("test", func(t *testing.T) {
			_, _ = postgres.SaveOwner(ctx, repoDTO.SaveOwnerDTO{
				Identifier: d.Identifier,
			})

			for _, f := range d.files {
				_, err := postgres.SaveFile(ctx, repoDTO.SaveFileReqDTO{
					Identifier: d.Identifier,
					File:       f,
					Client:     d.ClientId,
				})
				if err != nil {
				    t.Fatalf("Failed to save, error {%v}", err)
				}
			}

			err := postgres.DeleteFile(ctx, repoDTO.DeleteFileReqDTO{
				Identifier: repoDTO.Identifier{
					Username: d.Username,
					Folder:   d.Folder,
				},
				Client:  d.ClientId,
				FileName: d.FileName,
			})

			if err != nil {
				if !d.Err {
					t.Errorf("expected none, got error {%v} in create", err)
				}
			}

			err = postgres.RestoreFile(ctx, repoDTO.RestoreFileReqDTO{
				Identifier: repoDTO.Identifier{
					Username: d.Username,
					Folder:   d.Folder,
				},
				Client:  d.ClientId,
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