package repl

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/DXICIDE/MagisterLinguae/internal/database"
)

func (state *AppState) AnkiTest() error {
	words, err := state.Db.GetAnki(context.Background(), state.CurrentLanguage.ID)
	if err != nil {
		return err
	}
	fmt.Println("Here are the words u have not seen for a long time")
	for _, word := range words {
		fmt.Printf("Do you know the word: %s? y/n", word.TokenName)
		scanned := state.Scan(os.Stdin)
		if len(scanned) != 1 {
			fmt.Println("Dont understand the answer, skipping to next word")
			continue
		}
		lowerstr := strings.ToLower(scanned[0])

		//mark it if they wish to
		if lowerstr == "y" {
			updateWordSeenParams := database.UpdateWordSeenParams{TokenName: word.TokenName, LanguageID: state.CurrentLanguage.ID}
			state.Db.UpdateWordSeen(context.Background(), updateWordSeenParams)
			continue
		} else if lowerstr == "n" {
			continue
		}
	}
	return nil
}
