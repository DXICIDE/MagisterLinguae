package api

import (
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/DXICIDE/MagisterLinguae/internal/database"
)

func (h *Handler) Mark(w http.ResponseWriter, r *http.Request) {
	word := strings.ToLower(r.PathValue("word"))
	languageID := r.URL.Query().Get("language_id")

	ID, err := strconv.Atoi(languageID)
	if err != nil {
		http.Error(w, "invalid language_id", http.StatusBadRequest)
		return
	}

	wordDB, err := h.Db.GetWord(r.Context(), database.GetWordParams{TokenName: word, LanguageID: int32(ID)})

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, "word not found", http.StatusNotFound)
			return
		}
		log.Printf("GetWord error: %v", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	err = h.Db.UpdateWordKnown(r.Context(), database.UpdateWordKnownParams{TokenName: wordDB.TokenName, LanguageID: int32(ID)})

	if err != nil {
		log.Printf("UpdateWordKnown error: %v", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	wordDB.Known = true

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(WordResponse{
		Word:       wordDB.TokenName,
		Known:      wordDB.Known,
		Frequency:  int(wordDB.Frequency),
		LastSeenAt: wordDB.LastSeenAt,
	})
}
