package repl

import (
	"context"
	"fmt"
	"log"
	"os"
)

func Replfunc(state *AppState) {
	for {
		words := state.Scan(os.Stdin)
		if len(words) == 0 {
			Help()
			continue
		}
		switch words[0] {
		case "learn":
			sentence, err := state.Learn(words[1:])
			if err != nil {
				fmt.Printf("couldn't process the words: %s", err)
			}

			println(sentence)

			err = state.Markfrequency()
			if err != nil {
				fmt.Printf("couldn't prompt the user to mark the word: %s", err)
			}

		case "learnfile":
			sentence, err := state.LearnFromFileRepl(words[1:])
			if err != nil {
				fmt.Printf("couldn't process the file: %s", err)
			}

			println(sentence)

			err = state.Markfrequency()
			if err != nil {
				fmt.Printf("couldn't prompt the user to mark the word: %s", err)
			}

		case "mark":
			if len(words) > 1 {
				err := state.Mark(words[1:])
				if err != nil {
					fmt.Printf("couldn't mark the word as known: %s", err)
				}
			} else {
				fmt.Println("too few words!")
			}

		case "unmark":
			if len(words) > 1 {
				err := state.UnMark(words[1:])
				if err != nil {
					fmt.Printf("couldn't mark the word as unknown %s", err)
				}
			} else {
				fmt.Println("too few words!")
			}

		case "lookup":
			printstring, err := state.LookUpWord(words[1])
			if err != nil {
				fmt.Printf("couldn't lookup the word: %s", err)
			}
			fmt.Println(printstring)

		case "list":
			err := state.ListByFrequency()
			if err != nil {
				fmt.Printf("couldn't list the words: %s", err)
			}

		case "resetdb":
			//function made for testing/theoretically stays for user to use in case of wiping everything
			err := state.Db.ResetWords(context.Background())
			if err != nil {
				fmt.Printf("couldn't reset words: %s", err)
			}

			err = state.Db.ResetLanguages(context.Background())
			if err != nil {
				fmt.Printf("couldn't reset languages: %s", err)
			}

			err = state.Db.SeedLanguages(context.Background())
			if err != nil {
				log.Fatalf("couldn't seed languages: %s", err)
			}

			fmt.Println("db was successfully reset and reseeded!")

		case "q", "quit":
			fmt.Println("Hope u learned a lot, see you later!")
			fmt.Println("Goodbye!")
			return

		case "newlang":
			err := state.CustomLanguage()
			if err != nil {
				fmt.Printf("could not create custom languge: %s", err)
			}

		case "lang":
			fmt.Printf("Current language name: %s - %s\n", state.CurrentLanguage.Code, state.CurrentLanguage.Name)

		case "deletelang":
			err := state.DeletetLang()
			if err != nil {
				fmt.Printf("could not delete language: %s", err)
			}

		case "switch":
			language, err := state.SelectLang()
			if err != nil {
				fmt.Printf("could not select language: %s", err)
			}
			state.CurrentLanguage = language

		case "anki":
			err := state.AnkiTest()
			if err != nil {
				fmt.Printf("Anki test couldn't proceed: %s", err)
			}

		case "dict", "dictionary":
			err := state.Dictionary(words[1])
			if err != nil {
				fmt.Printf("dictionary failed: %s", err)
			}

		case "books":
			state.Books()

		case "help":
			Help()

		case "deephelp":
			DeepHelp()

		default:
			fmt.Println("I dont understand the command")

		}
	}
}
