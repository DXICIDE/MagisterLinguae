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
				log.Fatalf("couldn't process the words: %s", err)
			}

			println(sentence)

			err = state.Markfrequency()
			if err != nil {
				log.Fatal("couldn't prompt the user to mark the word")
			}

		case "learnfile":
			sentence, err := state.LearnFromFile(words[1:])
			if err != nil {
				log.Fatalf("couldn't process the file: %s", err)
			}

			println(sentence)

			err = state.Markfrequency()
			if err != nil {
				log.Fatal("couldn't prompt the user to mark the word")
			}

		case "mark":
			if len(words) > 1 {
				err := state.Mark(words[1:])
				if err != nil {
					log.Fatal("couldn't mark the word as known")
				}
			} else {
				fmt.Println("too few words!")
			}

		case "unmark":
			if len(words) > 1 {
				err := state.UnMark(words[1:])
				if err != nil {
					log.Fatal("couldn't mark the word as known")
				}
			} else {
				fmt.Println("too few words!")
			}

		case "lookup":
			printstring, err := state.LookUpWord(words[1])
			if err != nil {
				log.Fatal("couldn't lookup the word")
			}
			fmt.Println(printstring)

		case "list":
			err := state.ListByFrequency()
			if err != nil {
				log.Fatalf("couldn't list the words: %s", err)
			}

		case "resetdb":
			err := state.Db.ResetWords(context.Background())
			if err != nil {
				log.Fatalf("couldn't reset the db: %s", err)
			}
			fmt.Println("db was succesfully reset!")

		case "q", "quit":
			fmt.Println("Hope u learned a lot, see you later!")
			fmt.Println("Goodbye!")
			return

		case "switch":
			language, err := state.SelectLang()
			if err != nil {
				log.Fatalf("%s", err)
			}
			state.CurrentLanguage = language

		case "help":
			Help()

		case "deephelp":
			DeepHelp()

		default:
			fmt.Println("I dont understand the command")

		}
	}
}
