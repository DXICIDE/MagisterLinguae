package repl

import (
	"context"
	"database/sql"
	"strings"

	"github.com/DXICIDE/MagisterLinguae/internal/database"
)

func Mark(db *database.Queries, word string) error {
	word = strings.ToLower(word)
	wordDB, err := db.GetWord(context.Background(), word)
	if err == sql.ErrNoRows {
		createWord := database.CreateWordParams{TokenName: word, Known: true}
		_, err = db.CreateWord(context.Background(), createWord)
		return err
	}
	err = db.UpdateWordKnown(context.Background(), wordDB.TokenName)
	return err
}
