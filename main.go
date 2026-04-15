package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/DXICIDE/MagisterLinguae/internal/database"
	"github.com/DXICIDE/MagisterLinguae/internal/repl"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // registers the "postgres" driver
)

func main() {
	//loading .env
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: No .env file found, using system environment variables")
	}

	//loading DB
	dbURL := os.Getenv("DB_URL")
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		fmt.Printf("Couldn't connect to the DB: %s", err)
	}

	//loading db to Appstate, loads the db and the language
	dbQueries := database.New(db)
	state := &repl.AppState{}
	state.Db = dbQueries
	if os.Getenv("LAST_LANG") == "" {
		language, err := state.SelectLang()
		if err != nil {
			log.Fatalf("%s", err)
		}
		state.CurrentLanguage = language
	}
	language := os.Getenv("LAST_LANG")
	dbQueries.GetLanguageByCode(context.Background(), language)

	fmt.Printf("Current language used is %s, change the language using the command switch\n", state.CurrentLanguage.Name)
	fmt.Println("Good luck with learning!")

	//REPL main
	Replfunc(state)
}

func Replfunc(state *repl.AppState) {
	for {
		words := state.Scan(os.Stdin)
		if len(words) == 0 {
			repl.Help()
			return
		}
		switch words[0] {
		case "learn":
			sentence, err := state.Learn(words[1:])
			if err != nil {
				log.Fatal("couldn't process the words")
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
			repl.Help()

		case "deephelp":
			repl.DeepHelp()

		default:
			fmt.Println("I dont understand the command")

		}
	}
}
