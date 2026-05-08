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
	return mux
}
