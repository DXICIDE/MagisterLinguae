package api

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func (h *Handler) ResetDb(w http.ResponseWriter, r *http.Request) {
	type DestructResponse struct {
		Message      string `json:"message"`
		Wordsdeleted int    `json:"words_deleted"`
	}

	languageID := r.URL.Query().Get("language_id")
	ID, err := strconv.Atoi(languageID)
	if err != nil {
		http.Error(w, "invalid language_id", http.StatusBadRequest)
		return
	}

	words, err := h.Db.GetWords(r.Context(), int32(ID))
	if err != nil {
		http.Error(w, "Could not retrieve the words from db", http.StatusInternalServerError)
		return
	}

	err = h.Db.ResetLangWords(r.Context(), int32(ID))
	if err != nil {
		http.Error(w, "Could not reset the language", http.StatusInternalServerError)
		return
	}

	response := DestructResponse{Message: "database reset successfully", Wordsdeleted: len(words)}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
