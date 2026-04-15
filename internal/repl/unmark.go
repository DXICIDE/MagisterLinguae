package repl

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/DXICIDE/MagisterLinguae/internal/database"
)

func (state *AppState) UnMark(words []string) error {
	for _, v := range words {
		word := strings.ToLower(v)

		getWordParams := database.GetWordParams{TokenName: word, LanguageID: state.CurrentLanguage.ID}

		wordDB, err := state.Db.GetWord(context.Background(), getWordParams)
		if err == sql.ErrNoRows {
			fmt.Println("The word is already unknown")
			return nil
		} else if err != nil {
			return err
		}

		wordWordUnknown := database.UpdateWordUnKnownParams{TokenName: wordDB.TokenName, LanguageID: state.CurrentLanguage.ID}

		err = state.Db.UpdateWordUnKnown(context.Background(), wordWordUnknown)
		if err != nil {
			return nil
		}
	}
	return nil
}
