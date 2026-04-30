package repl

import (
	"context"
	"os"
	"testing"

	_ "github.com/lib/pq" // registers the "postgres" driver
)

func TestConfig(t *testing.T) {
	dbQueries := setupTestDB(t)

	state := &AppState{}
	state.Db = dbQueries

	var err error
	state.CurrentLanguage, err = state.Db.GetLanguageById(context.Background(), 1)
	if err != nil {
		t.Fatalf("Getting the language failed: %s", err)
	}
	config := Config{LastLanguage: 1}
	state.SaveConfig(config)
	defer os.Remove("conf.json")

	// Test
	config, err = state.GetConfig()
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	// Check output
	expected := "it"
	if config.LastLanguage != 1 {
		t.Fatalf("expected %q in output, got %q", expected, config.LastLanguage)
	}

}
