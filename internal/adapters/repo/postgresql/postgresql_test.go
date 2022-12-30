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

	truncate()
}

func TestRenameOwner(t *testing.T) {

	truncate()

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

	truncate()
}

func TestDeleteOwner(t *testing.T) {

	truncate()

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

	truncate()
}

func TestFindAllOwners(t *testing.T) {

	truncate()

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

	truncate()
}

// func TestDeleteOwner(t *testing.T) {

// 	truncate()

// 	data := []struct {
// 		Username string
// 		Folder   string
// 		Err      bool
// 	}{
// 		{
// 			Username: "u1",
// 			Folder:   "f1",
// 			Err:      false,
// 		},
// 	}

// 	for _, d := range data {
// 		t.Run("test", func(t *testing.T) {
// 			id, _ := postgres.SaveOwner(ctx, d.Username, d.Folder)
// 			id_, err := postgres.DeleteOwner(ctx, id)
// 			if d.Err {
// 				if err != nil {
// 					t.Errorf("expected none, got error")
// 				}
// 			}
// 			if id_ != id {
// 				t.Errorf("expected id=%d, got id=%d", id, id_)
// 			}
// 		})
// 	}

// 	_, err := postgres.conn.Exec(ctx, "TRUNCATE TABLE owners CASCADE")
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	truncate()
// }

// func TestUpdateOwner(t *testing.T) {

// 	truncate()

// 	data := []struct {
// 		Username  string
// 		Folder    string
// 		NewFolder string
// 		Err       bool
// 	}{
// 		{
// 			Username:  "u1",
// 			Folder:    "f1",
// 			NewFolder: "ff1",
// 			Err:       false,
// 		},
// 	}

// 	for _, d := range data {
// 		t.Run("test", func(t *testing.T) {
// 			id, _ := postgres.SaveOwner(ctx, d.Username, d.Folder)
// 			id_, err := postgres.RenameOwner(ctx, id, d.NewFolder)
// 			if d.Err {
// 				if err != nil {
// 					t.Errorf("expected none, got error")
// 				}
// 			}
// 			if id_ != id {
// 				t.Errorf("expected id=%d, got id=%d", id, id_)
// 			}
// 		})
// 	}

// 	_, err := postgres.conn.Exec(ctx, "TRUNCATE TABLE owners CASCADE")
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	truncate()
// }
