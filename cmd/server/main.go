package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/DXICIDE/MagisterLinguae/internal/api"
	"github.com/DXICIDE/MagisterLinguae/internal/database"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // registers the "postgres" driver
)

func main() {
	godotenv.Load()
	handler := &api.Handler{}
	dbURL := os.Getenv("DB_URL")
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("couldn't connect to the database: %v", err)
	}

	dbQueries := database.New(db)
	handler.Db = dbQueries

	mux := api.NewRouter(handler)

	s := http.Server{
		Handler: mux,
		Addr:    ":8080",
	}

	if err := s.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
