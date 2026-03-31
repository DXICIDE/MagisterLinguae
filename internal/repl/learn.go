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

	//for loop for checking each word individually
	for _, wordName := range words {
		fmt.Print(wordName)
		word, err := db.GetWord(ctx, wordName)

		if err == sql.ErrNoRows { //if the word is not in db
			fmt.Println("Word not found")
			createWord := database.CreateWordParams{TokenName: wordName, Known: false}
			db.CreateWord(ctx, createWord)
			proccesedWords = append(proccesedWords, wordName)
		} else if err != nil { //other errors
			return nil, fmt.Errorf("database error getting word '%s': %w", wordName, err)
		} else { //if the word is in db
			fmt.Printf("Word '%s' already exists\n", wordName)
			proccesedWords = append(proccesedWords, word.TokenName)
		}
	}
	return proccesedWords, nil
}
