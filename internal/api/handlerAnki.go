package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/DXICIDE/MagisterLinguae/internal/database"
)

func (h *Handler) Anki(w http.ResponseWriter, r *http.Request) {
	var anki struct {
		ID    int32 `json:"language_id"`
		Count int32 `json:"count"`
	}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&anki)
	if err != nil {
		log.Printf("Error decoding parameters: %s", err)
		http.Error(w, "wrong json format", http.StatusBadRequest)
		return
	}

	ankiWords, err := h.Db.GetAnki(r.Context(), database.GetAnkiParams{LanguageID: anki.ID, Limit: anki.Count})
	if err != nil {
		log.Printf("GetAnki error: %s", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	response := make([]WordResponse, 0, len(ankiWords))
	for _, w := range ankiWords {
		response = append(response, WordResponse{
			Word:       w.TokenName,
			Known:      w.Known,
			Frequency:  int(w.Frequency),
			LastSeenAt: w.LastSeenAt,
		})
	}

	json.NewEncoder(w).Encode(response)
}

func (h *Handler) AnkiAnswer(w http.ResponseWriter, r *http.Request) {
	var body struct {
		TokenName  string `json:"token_name"`
		LanguageID int32  `json:"language_id"`
		Known      bool   `json:"known"`
	}

	var response struct {
		Ok bool `json:"ok"`
	}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&body)
	if err != nil {
		log.Printf("Error decoding parameters: %s", err)
		http.Error(w, "wrong json format", http.StatusBadRequest)
		return
	}

	if body.Known {
		updateWordSeenParams := database.UpdateWordSeenParams{TokenName: body.TokenName, LanguageID: body.LanguageID}
		err := h.Db.UpdateWordSeen(r.Context(), updateWordSeenParams)
		if err != nil {
			log.Printf("UpdateWordSeen error: %s", err)
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}
	}

	response.Ok = true

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
