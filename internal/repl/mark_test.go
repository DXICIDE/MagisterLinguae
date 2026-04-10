package repl

import (
	"context"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	_ "github.com/lib/pq" // registers the "postgres" driver
)

func TestMark(t *testing.T) {
	tests := map[string]struct {
		input      string
		wordToMark []string
	}{
		"simple": {
			input:      "una",
			wordToMark: []string{"una"},
		},

		"multiple": {
			input:      "ciao io sono rostislav",
			wordToMark: []string{"Ciao", "io", "sono", "Rostislav"},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			//setup the db and scan the input
			dbQueries := setupTestDB(t)
			reader := strings.NewReader(tc.input)
			words := Scan(reader)

			//add known words, if there are any
			err := Mark(dbQueries, tc.wordToMark)
			if err != nil {
				t.Fatalf("setup failed")
			}

			for _, v := range words {
				word, err := dbQueries.GetWord(context.Background(), v)
				if err != nil {
					t.Fatalf("couldnt get word: %s err: %s", v, err)
				}
				if diff := cmp.Diff(true, word.Known); diff != "" {
					t.Fatalf("%s", diff)
				}
			}
		})
	}
}
