package repl

import (
	"context"
	"database/sql"
	"os"
	"testing"

	"github.com/DXICIDE/MagisterLinguae/internal/database"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // registers the "postgres" driver
)

func setupTestDB(t *testing.T) *database.Queries {
	t.Helper()
	err := godotenv.Load("../../.env")
	if err != nil {
		t.Logf("Warning: Could not load .env: %s", err)
	}

	dbURL := os.Getenv("TEST_DB_URL")

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		t.Fatalf("Couldn't connect to test DB: %s", err)
	}

	dbQueries := database.New(db)
	dbQueries.ResetWords(context.Background())

	t.Cleanup(func() {
		dbQueries.ResetWords(context.Background())
		db.Close()
	})

	return dbQueries
}
