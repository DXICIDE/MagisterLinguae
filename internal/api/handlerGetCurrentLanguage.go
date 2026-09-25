package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/DXICIDE/MagisterLinguae/internal/database"
	"github.com/DXICIDE/MagisterLinguae/internal/repl"
)

func (h *Handler) GetCurrentLanguage(w http.ResponseWriter, r *http.Request) {
	lastlang, err := repl.GetConfig()

	if lastlang == (repl.Config{}) {
		languagesList, err := h.Db.GetLanguageList(r.Context())
		if err != nil {
			log.Printf("languagesList error: %v", err)
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}

		if len(languagesList) > 0 {
			lastlang.LastLanguage = int32(languagesList[0].ID)
		}

		if len(languagesList) == 0 {
			lang, err := h.Db.CreateLanguage(r.Context(), database.CreateLanguageParams{Code: "en", Name: "English"})
			if err != nil {
				log.Printf("languagesList error: %v", err)
				http.Error(w, "internal server error", http.StatusInternalServerError)
				return
			}
			lastlang.LastLanguage = int32(lang.ID)
		}
		repl.SaveConfig(lastlang)
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
