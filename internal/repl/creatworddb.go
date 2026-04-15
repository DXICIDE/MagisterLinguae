package repl

import (
	"context"
	"time"

	"github.com/DXICIDE/MagisterLinguae/internal/database"
)

func (state *AppState) createWordDB(word string, knownbool bool) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	createWord := database.CreateWordParams{TokenName: word, Known: knownbool, LanguageID: state.CurrentLanguage.ID}
	_, err := state.Db.CreateWord(ctx, createWord)
	return err
}
