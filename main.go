package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/DXICIDE/MagisterLinguae/internal/database"
	"github.com/DXICIDE/MagisterLinguae/internal/repl"
	_ "github.com/lib/pq" // registers the "postgres" driver
)

type apiConfig struct {
	db *database.Queries
}

func main() {
	dbURL := os.Getenv("DB_URL")
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		fmt.Printf("Couldn't connect to the DB: %s", err)
	}
	dbQueries := database.New(db)
	apiCfg := &apiConfig{}
	apiCfg.db = dbQueries

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
