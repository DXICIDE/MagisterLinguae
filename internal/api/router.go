package api

import (
	"net/http"

	"github.com/DXICIDE/MagisterLinguae/internal/database"
)

type Handler struct {
	Db *database.Queries
}

func NewRouter(h *Handler) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /api/words/{word}", h.GetWord)
	mux.HandleFunc("GET /api/languages", h.GetLanguages)
	mux.HandleFunc("GET /api/languages/current", h.GetCurrentLanguage)
	mux.HandleFunc("PUT /api/languages/current", h.PutCurrentLanguage)
	mux.HandleFunc("PUT /api/words/{word}/mark", h.Mark)
	mux.HandleFunc("PUT /api/words/{word}/unmark", h.Unmark)
	mux.HandleFunc("POST /api/texts/process", h.ProcessText)
	mux.HandleFunc("GET /api/words", h.ListWords)
	mux.HandleFunc("GET /api/dictionary/{word}", h.Dictionary)
	mux.HandleFunc("POST /api/languages", h.InsertLanguage)
	mux.HandleFunc("DELETE /api/languages/{id}", h.DeleteLanguage)
	return mux
}
