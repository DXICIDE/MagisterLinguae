package api

import (
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/DXICIDE/MagisterLinguae/internal/repl"
)

func (h *Handler) PutCurrentLanguage(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	params := repl.Config{}
	err := decoder.Decode(&params)
	if err != nil {
		log.Printf("Error decoding parameters: %s", err)
		http.Error(w, "wrong json format", http.StatusBadRequest)
		return
	}

	language, err := h.Db.GetLanguageById(r.Context(), params.LastLanguage)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, "language not found", http.StatusNotFound)
			return
		}
		log.Printf("currentLanguage error: %v", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	err = repl.SaveConfig(repl.Config{LastLanguage: params.LastLanguage})
	if err != nil {
		log.Printf("Error saving config: %s", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(language)
}
