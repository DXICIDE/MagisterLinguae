package repl

import (
	"context"
	"strings"
	"testing"

	"github.com/DXICIDE/MagisterLinguae/internal/database"
	_ "github.com/lib/pq" // registers the "postgres" driver
)

func TestLookUpWord(t *testing.T) {
	dbQueries := setupTestDB(t)

	// Insert test word
	dbQueries.CreateWord(context.Background(), database.CreateWordParams{
		TokenName: "hello",
		Known:     true,
	})
	dbQueries.CreateWord(context.Background(), database.CreateWordParams{
		TokenName: "polo",
		Known:     false,
	})

	words := make([]string, 2)
	words[0] = "hello"
	Learn(dbQueries, words)
	words[1] = "polo"
	Learn(dbQueries, words)
	Learn(dbQueries, words)

	// Test
	outputhello, err := LookUpWord(dbQueries, "hello")
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	outputpolo, err := LookUpWord(dbQueries, "polo")
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	// Check output
	expected := "How many times you saw hello: 4\nDo you know word hello: Yes"
	if !strings.Contains(outputhello, expected) {
		t.Fatalf("expected %q in output, got %q", expected, outputhello)
	}

	expected = "How many times you saw polo: 3\nDo you know word polo: No"
	if !strings.Contains(outputpolo, expected) {
		t.Fatalf("expected %q in output, got %q", expected, outputpolo)
	}
}
