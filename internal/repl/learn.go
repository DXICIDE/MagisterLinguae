package repl

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/DXICIDE/MagisterLinguae/internal/database"
)

func Learn(db *database.Queries, words []string) ([]string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var proccesedWords []string

	//for loop for checking each word individually
	for _, wordName := range words {
		word := wordName
		word = strings.ToLower(word)
		if word == "," || word == "." || word == "?" || word == "!" {
			proccesedWords = append(proccesedWords, wordName)
			continue
		}

		wordDB, err := db.GetWord(ctx, word)
		if err == sql.ErrNoRows { //if the word is not in db
			fmt.Printf("Word %s not found\n", wordName)
			createWord := database.CreateWordParams{TokenName: word, Known: false}

			_, err = db.CreateWord(ctx, createWord)
			if err != nil {
				return nil, err
			}

			wordName = unknownWord(wordName)
			proccesedWords = append(proccesedWords, wordName)
		} else if err != nil { //other errors
			return nil, fmt.Errorf("database error getting word '%s': %w", wordName, err)

		} else { //if the word is in db
			fmt.Printf("Word '%s' already exists\n", wordName)
			if wordDB.Known == false {
				wordName = unknownWord(wordName)
			}
			proccesedWords = append(proccesedWords, wordName)
		}
	}
	return proccesedWords, nil
}

// adds a brackets if the word is unknown
func unknownWord(word string) string {
	return fmt.Sprintf("[%s]", word)
}
