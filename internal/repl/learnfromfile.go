package repl

import (
	"bufio"
	"os"
)

func (state *AppState) LearnFromFileRepl(filepath []string) (string, error) {
	var path string
	for _, v := range filepath {
		path += v
	}

	return state.LearnFromFilePath(path)
}

func (state *AppState) LearnFromFilePath(path string) (string, error) {
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
			words = state.tokenizeString(line)
			sentence, err = state.Learn(words)
			paragraphs += sentence
			break
		}
		words = state.tokenizeString(line)
		sentence, err = state.Learn(words)
		sentence = sentence + "\n"
		paragraphs += sentence
	}
	return paragraphs, nil
}
