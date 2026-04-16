package repl

import (
	"context"
	"errors"
	"fmt"
	"os"
)

// stdin scan
func (state *AppState) DeletetLang() error {

	//gets all languages for display
	languages, err := state.Db.GetLanguageList(context.Background())
	if err != nil {
		return err
	}

	for _, lang := range languages {
		fmt.Printf("%d - %s - %s\n", lang.ID, lang.Code, lang.Name)
	}
	fmt.Println("Type the code of language to delete please:")

	code := state.Scan(os.Stdin)
	if len(code) != 1 {
		return errors.New("Too many or too few words!")
	}

	err = state.Db.DeleteLanguage(context.Background(), code[0])
	if err != nil {
		return fmt.Errorf("couldnt delete language: %s", err)
	}

	fmt.Println("Language tab succesfully deleted!")
	return nil
}
