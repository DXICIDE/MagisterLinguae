package repl

import (
	"fmt"
	"strings"
)

func BackToSentence(proccesedWords []string) string {
	for i := 0; i < len(proccesedWords); i++ {
		if i+1 == len(proccesedWords) {
			continue
		}
		if proccesedWords[i+1] == "." || proccesedWords[i+1] == "," || proccesedWords[i+1] == "!" || proccesedWords[i+1] == "?" || proccesedWords[i+1] == "-" {
			continue
		}
		proccesedWords[i] = fmt.Sprintf("%s ", proccesedWords[i])
	}
	return strings.Join(proccesedWords, "")
}
