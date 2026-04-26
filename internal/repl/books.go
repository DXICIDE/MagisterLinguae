package repl

import (
	"fmt"
	"log"
	"os"
)

func (state *AppState) Books() {
	for {
		fmt.Println("Hello, welcome to the books section! Please choose a book by typing a number:")
		fmt.Println("1. L'italiano Secondo Il Metodo Natura")
		words := state.Scan(os.Stdin)
		if len(words) == 0 {
			fmt.Println("Too many or too few words! type q to exit")
			continue
		}
		switch words[0] {
		case "1", "1.":
			name := "L'italiano_Secondo_Il_Metodo_Natura"
			state.chapter(name)
			continue

		case "q", "quit":
			return
		}

	}
}

func (state *AppState) chapter(bookname string) {
	path := fmt.Sprintf("internal/repl/books/%s/", bookname)
	chapterCount, err := countSubdirs(path)
	if err != nil {
		fmt.Printf("Couldn't count the subcategories: %s\n", err)
		return
	}
	for {
		fmt.Printf("This book has %d chapters!\n", chapterCount)
		fmt.Println("Please choose a chapter number!:")
		words := state.Scan(os.Stdin)
		if len(words) == 0 {
			fmt.Println("Too many or too few words!")
			return
		}
		if words[0] == "q" {
			return
		}
		state.readPage(bookname, words[0])
	}
}

func (state *AppState) readPage(bookname string, chapterNmb string) {
	path := fmt.Sprintf("internal/repl/books/%s/chapter%s/", bookname, chapterNmb)
	chapterCount, err := countSubdirs(path)
	if err != nil {
		fmt.Println("Couldn't count the subcategories")
		return
	}
	for {
		fmt.Printf("This chapter has %d pages!\n", chapterCount)
		fmt.Println("Please choose a page number!:")
		words := state.Scan(os.Stdin)
		if len(words) == 0 {
			fmt.Println("Too many or too few words!")
			return
		}
		if words[0] == "q" {
			return
		}
		pageDir := path + fmt.Sprintf("page%s/", words[0])
		pathFile := pageDir + "page.txt"

		sentence, err := state.LearnFromFilePath(pathFile)
		if err != nil {
			log.Fatalf("couldn't process the file: %s", err)
		}
		println(sentence)
		err = state.Markfrequency()

		if err != nil {
			log.Fatal("couldn't prompt the user to mark the word")
		}
	}

}
func countSubdirs(dirPath string) (int, error) {
	entries, err := os.ReadDir(dirPath)
	if err != nil {
		return 0, err
	}
	count := 0
	for _, entry := range entries {
		if entry.IsDir() {
			count++
		}
	}
	return count, nil
}
