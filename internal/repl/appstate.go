package repl

import "github.com/DXICIDE/MagisterLinguae/internal/database"

type AppState struct {
	Db              *database.Queries
	CurrentLanguage database.Language
}
