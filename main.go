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

type apiConfig struct {
	db *database.Queries
}

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

	//loading db to apiconfig, its still a work in progress, just a reminder to make the web ui in the future
	dbQueries := database.New(db)
	apiCfg := &apiConfig{}
	apiCfg.db = dbQueries

	//REPL main
	replfunc(apiCfg)
}

func replfunc(apiCfg *apiConfig) {
	for {
		words := repl.Scan(os.Stdin)
		if len(words) == 0 {
			fmt.Println("Nothing was inputed")
		}

		switch words[0] {
		case "learn":
			processedWords, err := repl.Learn(apiCfg.db, words[1:])
			if err != nil {
				log.Fatal("couldn't process the words")
			}

			sentence := repl.BackToSentence(processedWords)
			println(sentence)

			err = repl.Markfrequency(apiCfg.db)
			if err != nil {
				log.Fatal("couldn't prompt the user to mark the word")
			}

		case "mark":
			if len(words) > 1 {
				repl.Mark(apiCfg.db, words[1:])
			} else {
				fmt.Println("too few words!")
			}

		case "lookup":
			printstring, err := repl.LookUpWord(apiCfg.db, words[1])
			if err != nil {
				log.Fatal("couldn't lookup the word")
			}
			fmt.Println(printstring)
		case "resetdb":
			err := apiCfg.db.ResetWords(context.Background())
			if err != nil {
				log.Fatal("couldn't reset the db")
			}
			fmt.Println("db was succesfully reset!")
		case "q", "quit":
			fmt.Println("Hope u learned a lot, see you later!")
			fmt.Println("Goodbye!")
			return

		default:
			fmt.Println("I dont understand the command")

		}
	}
}
