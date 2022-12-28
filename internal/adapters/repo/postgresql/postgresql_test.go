package postgresql

import (
	"context"
	"testing"
)

func TestFindAll(t *testing.T) {
	ctx := context.Background()
	url := "postgres://postgres:postgres@localhost:5431/postgres?sslmode=disable"
	connection, err := New(ctx, url)
	if err != nil {
		t.Fatal(err)
	}
	files, err := connection.FindAll(ctx, "u1", "f1")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(files)
}