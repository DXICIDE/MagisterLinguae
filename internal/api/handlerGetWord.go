package api

import (
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/DXICIDE/MagisterLinguae/internal/database"
)

type WordResponse struct {
	Word       string    `json:"word"`
	Known      bool      `json:"known"`
	Frequency  int       `json:"frequency"`
	LastSeenAt time.Time `json:"last_seen_at"`
}

func (h *Handler) GetWord(w http.ResponseWriter, r *http.Request) {
	word := r.PathValue("word")
	languageID := r.URL.Query().Get("language_id")

	ID, err := strconv.Atoi(languageID)
	if err != nil {
		http.Error(w, "invalid language_id", http.StatusBadRequest)
		return
	}

	wordDB, err := h.Db.GetWord(r.Context(), database.GetWordParams{
		TokenName:  word,
		LanguageID: int32(ID),
	})

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, "word not found", http.StatusNotFound)
			return
		}
		log.Printf("GetWord error: %v", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(WordResponse{
		Word:       wordDB.TokenName,
		Known:      wordDB.Known,
		Frequency:  int(wordDB.Frequency),
		LastSeenAt: wordDB.LastSeenAt,
	})

}
