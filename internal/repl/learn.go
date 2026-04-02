package repl

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/DXICIDE/MagisterLinguae/internal/database"
)

func Learn(db *database.Queries, words []string) ([]string, error) {

	var proccesedWords []string

	//for loop for checking each word individually
	for _, wordName := range words {
		//if the word is any of these, we dont do anything
		word := wordName
		word = strings.ToLower(word)
		if word == "," || word == "." || word == "?" || word == "!" {
			proccesedWords = append(proccesedWords, wordName)
			continue
		}

		wordDB, err := db.GetWord(context.Background(), word)

		//if the word is not in db
		if err == sql.ErrNoRows {
			fmt.Printf("Word %s not found\n", wordName)
			err = createWordDB(word, db, false)
			if err != nil {
				return nil, err
			}

			wordName = unknownWord(wordName)
			proccesedWords = append(proccesedWords, wordName)

			//other errors
		} else if err != nil {
			return nil, fmt.Errorf("database error getting word '%s': %w", wordName, err)

			//if the word is in db
		} else {
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
