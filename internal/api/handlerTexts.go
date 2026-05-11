package api

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/DXICIDE/MagisterLinguae/internal/database"
	"github.com/DXICIDE/MagisterLinguae/internal/repl"
)

type ProcessTextResponse struct {
	ProcessedText string    `json:"processed_text"`
	Stats         TextStats `json:"stats"`
}

type TextStats struct {
	TotalWords   int `json:"total_words"`
	KnownWords   int `json:"known_words"`
	UnknownWords int `json:"unknown_words"`
}

var body struct {
	Text       string `json:"text"`
	LanguageID int32  `json:"language_id"`
}

func (h *Handler) ProcessText(w http.ResponseWriter, r *http.Request) {
	var totalWords, knownWords, unknownWords int

	decoder := json.NewDecoder(r.Body)
	params := body
	err := decoder.Decode(&params)

	log.Printf("Received text: %q, language_id: %d", params.Text, params.LanguageID)
	if err != nil {
		http.Error(w, "invalid language_id", http.StatusBadRequest)
		return
	}

	lines := strings.Split(params.Text, "\n")
	var processedLines []string

	for _, line := range lines {
		var processedTokens []string
		words := repl.TokenizeString(line)
		for _, wordName := range words {
			//if the word is any of these, we dont do anything
			word := wordName
			word = strings.ToLower(word)
			if word == "," || word == "." || word == "?" || word == "!" || word == "-" || word == ";" || word == ":" || word == "«" || word == "»" || word == "(" || word == ")" {
				processedTokens = append(processedTokens, wordName)
				continue
			}

			getWordParams := database.GetWordParams{TokenName: word, LanguageID: params.LanguageID}
			wordDB, err := h.Db.GetWord(context.Background(), getWordParams)

			totalWords++
			if wordDB.Known {
				knownWords++
			} else {
				unknownWords++
			}

			//if the word is not in db
			if err == sql.ErrNoRows {
				err = h.createWordDB(word, false, params.LanguageID)
				if err != nil {
					log.Printf("Error creating word: %s", err)
					http.Error(w, "internal server error", http.StatusInternalServerError)
					return
				}

				wordName = unknownWord(wordName)
				processedTokens = append(processedTokens, wordName)

				//other errors
			} else if err != nil {
				log.Printf("database error getting word: %s", err)
				http.Error(w, "internal server error", http.StatusInternalServerError)
				return

				//if the word is in db
			} else {

				//update the last time we saw the word and frequency +1
				updateWordSeenParams := database.UpdateWordSeenParams{TokenName: word, LanguageID: params.LanguageID}
				h.Db.UpdateWordSeen(context.Background(), updateWordSeenParams)

				//add brackets
				if wordDB.Known == false {
					wordName = unknownWord(wordName)
				}

				processedTokens = append(processedTokens, wordName)
			}
		}

		processedLine := repl.BackToSentence(processedTokens)
		processedLines = append(processedLines, processedLine)
	}

	//for loop for checking each word individually

	result := strings.Join(processedLines, "\n")

	w.Header().Set("Content-Type", "application/json")
	response := ProcessTextResponse{
		ProcessedText: result,
		Stats: TextStats{
			TotalWords:   totalWords,
			KnownWords:   knownWords,
			UnknownWords: unknownWords,
		},
	}
	json.NewEncoder(w).Encode(response)
}

// adds a brackets if the word is unknown
func unknownWord(word string) string {
	return fmt.Sprintf("[%s]", word)
}

func (h *Handler) createWordDB(word string, knownbool bool, ID int32) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	createWord := database.CreateWordParams{TokenName: word, Known: knownbool, LanguageID: ID}
	_, err := h.Db.CreateWord(ctx, createWord)
	return err
}
