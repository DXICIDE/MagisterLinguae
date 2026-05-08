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

type UserConfig struct {
	LastLanguage int32 `json:"last_language"`
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
		log.Fatalf("Couldn't connect to the DB: %s", err)
	}

	//loading db to Appstate, loads the db and the language
	dbQueries := database.New(db)
	state := &repl.AppState{}
	state.Db = dbQueries

	//loading config
	userConfig, err := repl.GetConfig()
	if err != nil {
		log.Fatalf("could not get the user config: %s", err)
	}

	//loading last language used
	language := LastLanguage(userConfig, state)
	state.CurrentLanguage = language

	fmt.Printf("Current language used is %s, change the language using the command switch\n", state.CurrentLanguage.Name)
	fmt.Println("Good luck with learning!")

	//REPL main
	repl.Replfunc(state)
}

func LastLanguage(userConfig repl.Config, state *repl.AppState) database.Language {
	if userConfig.LastLanguage == 0 {
		language, err := state.SelectLang()
		if err != nil {
			log.Fatalf("%s", err)
		}
		return language
	} else {
		language, err := state.Db.GetLanguageById(context.Background(), userConfig.LastLanguage)
		if err != nil {
			log.Fatalf("%s", err)
		}
		return language
	}
}
