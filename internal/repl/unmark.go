package repl

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
)

func (state *AppState) UnMark(words []string) error {
	for _, v := range words {
		word := strings.ToLower(v)

		wordDB, err := state.Db.GetWord(context.Background(), word)
		if err == sql.ErrNoRows {
			fmt.Println("The word is already unknown")
			return nil
		} else if err != nil {
			return err
		}

		err = state.Db.UpdateWordUnKnown(context.Background(), wordDB.TokenName)
		if err != nil {
			return nil
		}
	}
	return nil
}
