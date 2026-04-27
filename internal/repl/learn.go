package repl

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/DXICIDE/MagisterLinguae/internal/database"
)

func (state *AppState) Learn(words []string) (string, error) {

	var proccesedWords []string

	//for loop for checking each word individually
	for _, wordName := range words {
		//if the word is any of these, we dont do anything
		word := wordName
		word = strings.ToLower(word)
		if word == "," || word == "." || word == "?" || word == "!" || word == "-" || word == ";" || word == ":" || word == "«" || word == "»" || word == "(" || word == ")" {
			proccesedWords = append(proccesedWords, wordName)
			continue
		}

		getWordParams := database.GetWordParams{TokenName: word, LanguageID: state.CurrentLanguage.ID}
		wordDB, err := state.Db.GetWord(context.Background(), getWordParams)

		//if the word is not in db
		if err == sql.ErrNoRows {
			fmt.Printf("Word %s not found\n", wordName)
			err = state.createWordDB(word, false)
			if err != nil {
				return "", err
			}

			wordName = unknownWord(wordName)
			proccesedWords = append(proccesedWords, wordName)

			//other errors
		} else if err != nil {
			return "", fmt.Errorf("database error getting word '%s': %w", wordName, err)

			//if the word is in db
		} else {
			fmt.Printf("Word '%s' already exists\n", word)

			//update the last time we saw the word and frequency +1
			updateWordSeenParams := database.UpdateWordSeenParams{TokenName: word, LanguageID: state.CurrentLanguage.ID}
			state.Db.UpdateWordSeen(context.Background(), updateWordSeenParams)

			//add brackets
			if wordDB.Known == false {
				wordName = unknownWord(wordName)
			}

			proccesedWords = append(proccesedWords, wordName)
		}
	}
	sentence := BackToSentence(proccesedWords)

	return sentence, nil
}

// adds a brackets if the word is unknown
func unknownWord(word string) string {
	return fmt.Sprintf("[%s]", word)
}
