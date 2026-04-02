package repl

import (
	"context"
	"database/sql"
	"strings"
	"time"

	"github.com/DXICIDE/MagisterLinguae/internal/database"
)

func Mark(db *database.Queries, word string) error {
	word = strings.ToLower(word)

	wordDB, err := db.GetWord(context.Background(), word)
	if err == sql.ErrNoRows {
		err = createWordDB(word, db, true)
		return err
	}

	err = db.UpdateWordKnown(context.Background(), wordDB.TokenName)

	return err
}

func createWordDB(word string, db *database.Queries, knownbool bool) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	createWord := database.CreateWordParams{TokenName: word, Known: knownbool}
	_, err := db.CreateWord(ctx, createWord)
	return err
}
