package repl

import (
	"context"
	"database/sql"
	"strings"

	"github.com/DXICIDE/MagisterLinguae/internal/database"
)

func (state *AppState) Mark(words []string) error {
	for _, v := range words {
		word := strings.ToLower(v)

		wordParams := database.GetWordParams{TokenName: word, LanguageID: state.CurrentLanguage.ID}

		wordDB, err := state.Db.GetWord(context.Background(), wordParams)
		if err == sql.ErrNoRows {
			err = state.createWordDB(word, true)
			if err != nil {
				return nil
			}
		}

		updateWordKnownParams := database.UpdateWordKnownParams{TokenName: wordDB.TokenName, LanguageID: state.CurrentLanguage.ID}

		err = state.Db.UpdateWordKnown(context.Background(), updateWordKnownParams)
		if err != nil {
			return nil
		}
	}
	return nil
}
