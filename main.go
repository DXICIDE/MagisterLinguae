package main

import (
	"fmt"

	"github.com/DXICIDE/MagisterLinguae/internal/repl"
)

func main() {
	for true {
		words := repl.Scan()
		if len(words) == 0 {
			continue
		}

		switch words[0] {
		case "known":

		case "scan":

		case "q":
			fmt.Println("exiting")
			return
		default:
			fmt.Println("I dont understand the command")
		}
	}
}
