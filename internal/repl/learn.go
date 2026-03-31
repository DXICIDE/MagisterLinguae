package repl

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/DXICIDE/MagisterLinguae/internal/database"
)

func Learn(db *database.Queries, words []string) ([]string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var proccesedWords []string

	for _, n := range words {
		word, err := db.GetWord(ctx, n)
		if err == sql.ErrNoRows {
			fmt.Println("Word not found")
			createWord := database.CreateWordParams{TokenName: n, Known: false}
			db.CreateWord(ctx, createWord)
			proccesedWords = append(proccesedWords, n)
		} else if err != nil {
			return nil, err
		}
		proccesedWords = append(proccesedWords, word.TokenName)

	}
	return proccesedWords, nil
}
