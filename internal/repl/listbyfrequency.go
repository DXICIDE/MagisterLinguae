package repl

import (
	"context"
	"fmt"
)

func (state *AppState) ListByFrequency() error {
	words, err := state.Db.GetListByFrequency(context.Background(), state.CurrentLanguage.ID)
	if err != nil {
		return err
	}
	fmt.Println("Minimum times seen is 2")
	fmt.Println("Here are the words ranked from 1 in descending order")
	for i, v := range words {
		numb := i + 1
		fmt.Printf("%d) %s - Seen:%d Known:%t\n", numb, v.TokenName, v.Frequency, v.Known)
	}
	return nil
}
