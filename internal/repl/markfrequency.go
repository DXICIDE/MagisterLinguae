package repl

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/DXICIDE/MagisterLinguae/internal/database"
)

//user will be prompted to mark the word after the 8th time of seeing the word, then at 12th time and then never.

func (state *AppState) Markfrequency() error {
	//Get words that were mentioned more than 7 times and are unknown and user was not yet prompted
	words, err := state.Db.MarkWordsByFrequency(context.Background(), state.CurrentLanguage.ID)
	if err != nil {
		return err
	}
	for _, word := range words {
		for {
			//set prompted to true
			updateWordPromptedParams := database.UpdateWordPromptedParams{TokenName: word.TokenName, LanguageID: state.CurrentLanguage.ID}
			state.Db.UpdateWordPrompted(context.Background(), updateWordPromptedParams)

			fmt.Printf("Do you wish to mark word: %s as known?(Y/N)\n", word.TokenName)

			//wait for input
			scanned := state.Scan(os.Stdin)
			if len(scanned) != 1 {
				fmt.Println("Too many or too few words!")
				continue
			}
			lowerstr := strings.ToLower(scanned[0])

			//mark it if they wish to
			if lowerstr == "y" {
				markword := make([]string, 1)
				markword[0] = word.TokenName
				state.Mark(markword)
				break
			} else if lowerstr == "n" {
				break
			}
			fmt.Println("I dont understand the command!")
		}
	}
	return nil
}
