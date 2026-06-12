package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/DXICIDE/MagisterLinguae/internal/database"
)

func (h *Handler) InsertLanguage(w http.ResponseWriter, r *http.Request) {
	var newLanguage struct {
		Code string `json:"code"`
		Name string `json:"name"`
	}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newLanguage)
	if err != nil {
		log.Printf("Error decoding parameters: %s", err)
		http.Error(w, "wrong json format", http.StatusBadRequest)
		return
	}

	language, err := h.Db.CreateLanguage(r.Context(), database.CreateLanguageParams{Code: newLanguage.Code, Name: newLanguage.Name})
	if err != nil {
		log.Printf("CreateLanguage error:: %s", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(language)
}

func (h *Handler) DeleteLanguage(w http.ResponseWriter, r *http.Request) {
	languageID := r.PathValue("id")

	ID, err := strconv.Atoi(languageID)
	if err != nil {
		http.Error(w, "invalid language_id", http.StatusBadRequest)
		return
	}

	err = h.Db.DeleteLanguage(r.Context(), int32(ID))
	if err != nil {
		log.Printf("DeleteLanguage error:: %s", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	var response struct {
		Ok bool `json:"ok"`
	}

	response.Ok = true

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
