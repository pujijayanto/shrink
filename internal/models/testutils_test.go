package models

import (
	"context"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
)

func newTestDB(t *testing.T) *pgxpool.Pool {
	dsn := "postgres://postgres:admin@localhost:5432/shrink_test?sslmode=disable"

	dbPool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		t.Fatal(err)
	}

	// Read and execute the setup SQL script
	setupScript, err := os.ReadFile("./testdata/setup.sql")
	if err != nil {
		dbPool.Close()
		t.Fatal(err)
	}

	_, err = dbPool.Exec(context.Background(), string(setupScript))
	if err != nil {
		dbPool.Close()
		t.Fatal(err)
	}

	// Register a cleanup function to run the teardown script
	t.Cleanup(func() {
		defer dbPool.Close()

		teardownScript, err := os.ReadFile("./testdata/teardown.sql")
		if err != nil {
			t.Fatal(err)
		}

		_, err = dbPool.Exec(context.Background(), string(teardownScript))
		if err != nil {
			t.Fatal(err)
		}
	})

	return dbPool
}
