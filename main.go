package main

import (
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
		_, err = repl.Learn(apiCfg.db, words[1:])
	case "mark":

	case "lookup":

	default:
		fmt.Println("I dont understand the command")

	}
}
