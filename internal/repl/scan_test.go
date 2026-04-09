package repl

import (
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestScan(t *testing.T) {
	tests := map[string]struct {
		input string
		want  []string
	}{
		"simple":     {input: "una ristorante", want: []string{"una", "ristorante"}},
		"nothing":    {input: "", want: nil},
		"apostrophe": {input: "C'e una ristorante", want: []string{"C'e", "una", "ristorante"}},
		"multiple sentences": {input: "Ciao, mi chiamo Carlo e ho 18 anni. Oggi vorrei parlarvi della mia tipica giornata.",
			want: []string{"Ciao", ",", "mi", "chiamo", "Carlo", "e", "ho", "18", "anni", ".", "Oggi", "vorrei", "parlarvi", "della", "mia", "tipica", "giornata", "."}},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			reader := strings.NewReader(tc.input)
			got := Scan(reader)
			diff := cmp.Diff(tc.want, got)
			if diff != "" {
				t.Fatalf("%s", diff)
			}
		})
	}
}
