package repl

import (
	"context"
	"strings"
	"testing"

	"github.com/DXICIDE/MagisterLinguae/internal/database"
	"github.com/google/go-cmp/cmp"
	_ "github.com/lib/pq" // registers the "postgres" driver
)

func TestMark(t *testing.T) {
	tests := map[string]struct {
		input        string
		wordToMark   []string
		languageCode string
	}{
		"simple": {
			input:        "una",
			wordToMark:   []string{"una"},
			languageCode: "it",
		},

		"multiple": {
			input:        "ciao io sono rostislav",
			wordToMark:   []string{"Ciao", "io", "sono", "Rostislav"},
			languageCode: "it",
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			//setup the db and scan the input
			dbQueries := setupTestDB(t)
			reader := strings.NewReader(tc.input)
			state := &AppState{}
			state.Db = dbQueries

			var err error
			state.CurrentLanguage, err = state.Db.GetLanguageByCode(context.Background(), tc.languageCode)
			if err != nil {
				t.Fatalf("Getting the language failed: %s", err)
			}

			words := state.Scan(reader)

			//add known words, if there are any
			err = state.Mark(tc.wordToMark)
			if err != nil {
				t.Fatalf("setup failed")
			}

			for _, v := range words {
				getWordParams := database.GetWordParams{TokenName: v, LanguageID: state.CurrentLanguage.ID}
				word, err := dbQueries.GetWord(context.Background(), getWordParams)
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
