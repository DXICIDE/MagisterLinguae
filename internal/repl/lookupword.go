package repl

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/DXICIDE/MagisterLinguae/internal/database"
)

func LookUpWord(db *database.Queries, word string) error {
	word = strings.ToLower(word)

	wordDB, err := db.GetWord(context.Background(), word)
	if err == sql.ErrNoRows {
		fmt.Printf("Word %s not found\n", word)
		return nil
	} else if err != nil {
		return err
	}

	//stdout print
	fmt.Printf("How many times you saw %s: %d\n", word, wordDB.Frequency)
	if wordDB.Known == true {
		fmt.Printf("Do you know word %s: Yes\n", word)
	} else {
		fmt.Printf("Do you know word %s: No\n", word)
	}

	fmt.Printf("Last time you saw %s: ", word)
	fmt.Println(wordDB.LastSeenAt.Format(time.RFC822))
	return nil
}
