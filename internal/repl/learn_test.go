package repl

import (
	"context"
	"strings"
	"testing"

	"github.com/DXICIDE/MagisterLinguae/internal/database"
	"github.com/google/go-cmp/cmp"
	_ "github.com/lib/pq" // registers the "postgres" driver
)

func TestLearn(t *testing.T) {
	tests := map[string]struct {
		input        string
		knownWords   []string
		wantWords    string
		wantInDB     []string
		languageCode string
	}{
		"simple": {
			input:        "una ristorante",
			wantWords:    "[una] [ristorante]",
			wantInDB:     []string{"una", "ristorante"},
			languageCode: "it",
		},
		"known": {
			input:        "una scuola",
			knownWords:   []string{"scuola"},
			wantWords:    "[una] scuola",
			wantInDB:     []string{"una", "scuola"},
			languageCode: "it",
		},
		"combination": {
			input:        "Buongiorno! Oggi vi presento la mia famiglia. C'e vi",
			knownWords:   []string{"oggi", "presento"},
			wantWords:    "[Buongiorno]! Oggi [vi] presento [la] [mia] [famiglia]. [C'e] [vi]",
			wantInDB:     []string{"buongiorno", "oggi", "vi", "presento", "la", "mia", "famiglia", "c'e"},
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
			for _, w := range tc.knownWords {
				err := state.createWordDB(w, true)
				if err != nil {
					t.Fatalf("setup failed")
				}
			}

			//testing the learn
			got, err := state.Learn(words)
			if err != nil {
				t.Fatalf("Learn failed: %s", err)
			}

			//diif in the return of learn
			if diff := cmp.Diff(tc.wantWords, got); diff != "" {
				t.Fatalf("%s", diff)
			}

			//diff in the db
			for _, expectedWord := range tc.wantInDB {
				getWordParams := database.GetWordParams{TokenName: expectedWord, LanguageID: state.CurrentLanguage.ID}
				saved, err := dbQueries.GetWord(context.Background(), getWordParams)
				if err != nil {
					t.Fatalf("Word %s not found in DB: %s", expectedWord, err)
				}
				if saved.TokenName != expectedWord {
					t.Fatalf("Expected %s, got %s", expectedWord, saved.TokenName)
				}
			}
		})
	}
}
