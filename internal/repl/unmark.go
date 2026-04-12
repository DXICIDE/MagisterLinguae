package repl

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/DXICIDE/MagisterLinguae/internal/database"
)

func UnMark(db *database.Queries, words []string) error {
	for _, v := range words {
		word := strings.ToLower(v)

		wordDB, err := db.GetWord(context.Background(), word)
		if err == sql.ErrNoRows {
			fmt.Println("The word is already unknown")
			return nil
		} else if err != nil {
			return err
		}
		err = db.UpdateWordUnKnown(context.Background(), wordDB.TokenName)
		if err != nil {
			return nil
		}
	}
	return nil
}
