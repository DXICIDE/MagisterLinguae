package repl

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/DXICIDE/MagisterLinguae/internal/database"
)

func (state *AppState) LookUpWord(word string) (string, error) {
	word = strings.ToLower(word)

	getWordParams := database.GetWordParams{TokenName: word, LanguageID: state.CurrentLanguage.ID}
	wordDB, err := state.Db.GetWord(context.Background(), getWordParams)
	if err == sql.ErrNoRows {
		return fmt.Sprintf("Word %s not found\n", word), nil
	} else if err != nil {
		return "", err
	}

	known := "No"
	if wordDB.Known {
		known = "Yes"
	}

	//stdout print
	result := fmt.Sprintf("How many times you saw %s: %d\n", word, wordDB.Frequency)
	result += fmt.Sprintf("Do you know word %s: %s\n", word, known)
	result += fmt.Sprintf("Last time you saw %s: %s", word, wordDB.LastSeenAt.Format(time.RFC822))

	return result, nil
}
