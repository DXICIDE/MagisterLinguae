package repl

import (
	"bufio"
	"os"

	"github.com/DXICIDE/MagisterLinguae/internal/database"
)

func LearnFromFile(db *database.Queries, filepath []string) (string, error) {
	var path string
	for _, v := range filepath {
		path += v
	}
	println(path)

	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()
	r := bufio.NewReader(file)

	var words []string
	var sentence string
	var paragraphs string

	for {
		line, err := r.ReadString('\n')
		if err != nil {
			words = tokenizeString(line)
			sentence, err = Learn(db, words)
			paragraphs += sentence
			break
		}
		words = tokenizeString(line)
		sentence, err = Learn(db, words)
		sentence = sentence + "\n"
		paragraphs += sentence
	}
	return paragraphs, nil
}
