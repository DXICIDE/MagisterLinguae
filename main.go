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

	dbQueries := database.New(db)
	apiCfg := &apiConfig{}
	apiCfg.db = dbQueries

	//CLI main
	words := repl.Scan()
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
	case "mark":
		if len(words) == 2 {
			repl.Mark(apiCfg.db, words[1])
		} else {
			fmt.Println("too many or too few words!")
		}
	case "lookup":

	case "resetdb":
		err := apiCfg.db.ResetWords(context.Background())
		if err != nil {
			log.Fatal("couldn't reset the db")
		}
		fmt.Println("db was succesfully reset!")
	default:
		fmt.Println("I dont understand the command")

	}
}
