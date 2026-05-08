package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/DXICIDE/MagisterLinguae/internal/repl"
)

func (h *Handler) GetCurrentLanguage(w http.ResponseWriter, r *http.Request) {
	lastlang, err := repl.GetConfig()

	if lastlang == (repl.Config{}) {
		http.Error(w, "internal server error", http.StatusNotFound)
		return
	}

	if err != nil {
		log.Printf("currentLanguage error: %v", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	language, err := h.Db.GetLanguageById(r.Context(), lastlang.LastLanguage)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(language)
}
