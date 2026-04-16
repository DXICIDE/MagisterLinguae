package repl

import (
	"context"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestScan(t *testing.T) {
	tests := map[string]struct {
		input        string
		want         []string
		languageCode string
	}{
		"simple":     {input: "una ristorante", want: []string{"una", "ristorante"}, languageCode: "it"},
		"nothing":    {input: "", want: nil, languageCode: "it"},
		"apostrophe": {input: "C'e una ristorante", want: []string{"C'e", "una", "ristorante"}, languageCode: "it"},
		"multiple sentences": {input: "Ciao, mi chiamo Carlo e ho 18 anni. Oggi vorrei parlarvi della mia tipica giornata.",
			want: []string{"Ciao", ",", "mi", "chiamo", "Carlo", "e", "ho", "18", "anni", ".", "Oggi", "vorrei", "parlarvi", "della", "mia", "tipica", "giornata", "."}, languageCode: "it"},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			reader := strings.NewReader(tc.input)

			dbQueries := setupTestDB(t)
			state := &AppState{}
			state.Db = dbQueries

			var err error
			state.CurrentLanguage, err = state.Db.GetLanguageByCode(context.Background(), tc.languageCode)
			if err != nil {
				t.Fatalf("Getting the language failed: %s", err)
			}

			got := state.Scan(reader)
			diff := cmp.Diff(tc.want, got)
			if diff != "" {
				t.Fatalf("%s", diff)
			}
		})
	}
}
