package repl

import (
	"context"
	"fmt"
	"os"

	"github.com/DXICIDE/MagisterLinguae/internal/database"
)

func (state *AppState) CustomLanguage() error {
	fmt.Println("Please enter a code which u will use as a switch command, the code must be shorter than 11 characters")
	fmt.Println("I recommend using the short form ISO 639-1 standard aka the two letter abbreviations,")
	fmt.Println("otherwise the dictionary function will not work:")
	code := state.Scan(os.Stdin)
	if len(code) != 1 {
		return fmt.Errorf("Too many or too few words!")
	}

	fmt.Println("Please enter a name of the language tab:")
	name := state.Scan(os.Stdin)

	if len(name) != 1 {
		return fmt.Errorf("Too many or too few words!")
	}
	createLanguageParams := database.CreateLanguageParams{Code: code[0], Name: name[0]}
	state.Db.CreateLanguage(context.Background(), createLanguageParams)
	fmt.Println("New language tab was succesfully created!")
	fmt.Println("Change language using the command switch!")
	return nil
}
