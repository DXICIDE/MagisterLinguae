package repl

import (
	"context"
	"database/sql"
	"strings"
	"time"

	"github.com/DXICIDE/MagisterLinguae/internal/database"
)

func Mark(db *database.Queries, words []string) error {
	for _, v := range words {
		word := strings.ToLower(v)

		wordDB, err := db.GetWord(context.Background(), word)
		if err == sql.ErrNoRows {
			err = createWordDB(word, db, true)
			if err != nil {
				return nil
			}
		}

		err = db.UpdateWordKnown(context.Background(), wordDB.TokenName)
		if err != nil {
			return nil
		}
	}
	return nil
}

func createWordDB(word string, db *database.Queries, knownbool bool) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	createWord := database.CreateWordParams{TokenName: word, Known: knownbool}
	_, err := db.CreateWord(ctx, createWord)
	return err
}
