package api

import (
	"encoding/json"
	"log"
	"net/http"
)

func (h *Handler) GetLanguages(w http.ResponseWriter, r *http.Request) {
	languagesList, err := h.Db.GetLanguageList(r.Context())

	if err != nil {
		log.Printf("languagesList error: %v", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(languagesList)
}
