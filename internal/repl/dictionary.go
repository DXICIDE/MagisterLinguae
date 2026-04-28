package repl

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type Entry struct {
	Word     string    `json:"word"`
	Meanings []Meaning `json:"meanings"`
}

type Meaning struct {
	PartOfSpeech string       `json:"partOfSpeech"`
	Definitions  []Definition `json:"definitions"`
}

type Definition struct {
	Definition string `json:"definition"`
	Example    string `json:"example"`
}

// stdin scan
func (state *AppState) Dictionary(word string) error {

	httpAdress := fmt.Sprintf("https://api.dictionaryapi.dev/api/v2/entries/en/%s", word)

	response, err := http.Get(httpAdress)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(responseData))
	return nil
}
