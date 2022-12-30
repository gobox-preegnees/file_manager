package postgresql

import (
	"context"
	"os"
	// "strconv"
	"testing"
	// repoDTO "github.com/gobox-preegnees/file_manager/internal/adapters/repo"
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
		ID       int
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
			Folder:   "f1",
			Err:      false,
		},
	}

	for _, d := range data {
		t.Run("test", func(t *testing.T) {
			id, err := postgres.SaveOwner(ctx, d.Username, d.Folder)
			if d.Err {
				if err == nil {
					t.Errorf("expected error, got none")
				}
			}
		    t.Logf("id=%d", id)
		})
	}

	truncate()
}

func TestDeleteOwner(t *testing.T) {

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
	}

	for _, d := range data {
		t.Run("test", func(t *testing.T) {
			id, _ := postgres.SaveOwner(ctx, d.Username, d.Folder)
			id_, err := postgres.DeleteOwner(ctx, id)
			if d.Err {
				if err != nil {
					t.Errorf("expected none, got error")
				}
			}
			if id_ != id {
				t.Errorf("expected id=%d, got id=%d", id, id_)
			}
		})
	}

	_, err := postgres.conn.Exec(ctx, "TRUNCATE TABLE owners CASCADE")
	if err != nil {
		t.Fatal(err)
	}

	truncate()
}

func TestUpdateOwner(t *testing.T) {

	truncate()

	data := []struct {
		Username  string
		Folder    string
		NewFolder string
		Err       bool
	}{
		{
			Username:  "u1",
			Folder:    "f1",
			NewFolder: "ff1",
			Err:       false,
		},
	}

	for _, d := range data {
		t.Run("test", func(t *testing.T) {
			id, _ := postgres.SaveOwner(ctx, d.Username, d.Folder)
			id_, err := postgres.RenameOwner(ctx, id, d.NewFolder)
			if d.Err {
				if err != nil {
					t.Errorf("expected none, got error")
				}
			}
			if id_ != id {
				t.Errorf("expected id=%d, got id=%d", id, id_)
			}
		})
	}

	_, err := postgres.conn.Exec(ctx, "TRUNCATE TABLE owners CASCADE")
	if err != nil {
		t.Fatal(err)
	}

	truncate()
}