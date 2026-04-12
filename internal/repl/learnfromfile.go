package repl

import (
	"bufio"
	"fmt"
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

	for {
		line, err := r.ReadString('\n')
		if err != nil {
			break
		}

		fmt.Print(line)
	}
	return "", nil
}
