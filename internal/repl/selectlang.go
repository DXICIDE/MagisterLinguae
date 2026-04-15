package repl

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/DXICIDE/MagisterLinguae/internal/database"
	"github.com/joho/godotenv"
)

// stdin scan
func (state *AppState) SelectLang() (database.Language, error) {
	fmt.Println("Select language to learn please:")

	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: No .env file found, using system environment variables")
	}

	//gets all languages for display
	languages, err := state.Db.GetLanguageList(context.Background())
	if err != nil {
		return database.Language{}, err
	}
	for _, lang := range languages {
		fmt.Printf("%d - %s - %s\n", lang.ID, lang.Code, lang.Name)
	}
	fmt.Println("Select language by typing its code or number:")

	//scan the user input
	scanned := state.Scan(os.Stdin)
	if len(scanned) != 1 {
		fmt.Println("Too many or too few words!")
	}
	var language database.Language

	//checks if its id, if not then search for code
	if id, err := strconv.Atoi(scanned[0]); err == nil {
		language, err = state.Db.GetLanguageById(context.Background(), int32(id))
		if err != nil {
			return database.Language{}, fmt.Errorf("language with ID %d not found: %w", id, err)
		}
	} else {
		language, err = state.Db.GetLanguageByCode(context.Background(), scanned[0])
		if err != nil {
			return database.Language{}, fmt.Errorf("language with code '%s' not found: %w", scanned[0], err)
		}
	}

	fmt.Printf("Language was succesfully switched to %s\n", language.Name)

	return language, nil
}
