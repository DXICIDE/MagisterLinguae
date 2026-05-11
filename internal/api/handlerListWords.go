package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

func (h *Handler) ListWords(w http.ResponseWriter, r *http.Request) {
	languageID := r.URL.Query().Get("language_id")
	ID, err := strconv.Atoi(languageID)
	if err != nil {
		http.Error(w, "invalid language_id", http.StatusBadRequest)
		return
	}

	words, err := h.Db.GetListByFrequency(r.Context(), int32(ID))
	if err != nil {
		log.Printf("GetListByFrequency error: %v", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	response := make([]WordResponse, 0, len(words))
	for _, w := range words {
		response = append(response, WordResponse{
			Word:       w.TokenName,
			Known:      w.Known,
			Frequency:  int(w.Frequency),
			LastSeenAt: w.LastSeenAt,
		})
	}

	json.NewEncoder(w).Encode(response)

}
