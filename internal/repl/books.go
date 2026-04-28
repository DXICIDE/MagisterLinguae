package repl

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
)

func (state *AppState) Books() {
	for {
		fmt.Println("Hello, welcome to the books section! Please choose a book by typing a number:")
		fmt.Println("1. L'italiano Secondo Il Metodo Natura")
		words := state.Scan(os.Stdin)
		if len(words) == 0 {
			fmt.Println("No input!")
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
		fmt.Printf("\nThis book has %d chapters!\n", chapterCount)
		fmt.Println("Please choose a chapter number:")
		words := state.Scan(os.Stdin)
		if len(words) == 0 {
			fmt.Println("No input!")
			return
		}
		if words[0] == "q" {
			return
		}

		id, err := strconv.Atoi(words[0])
		if err != nil {
			fmt.Printf("This is not a number nor valid command!: %s", err)
			continue
		}
		if id > chapterCount || id < 1 {
			fmt.Println("This page does not exist!")
			continue
		}

		state.readPage(bookname, words[0])
	}
}

func (state *AppState) readPage(bookname string, chapterNmb string) {
	path := fmt.Sprintf("internal/repl/books/%s/chapter%s/", bookname, chapterNmb)
	pageCount, err := countSubdirs(path)
	if err != nil {
		fmt.Println("Couldn't count the subcategories")
		return
	}
	for {
		fmt.Printf("\nThis chapter has %d pages!\n", pageCount)
		fmt.Println("Please choose a page number or use dictionary!:")
		words := state.Scan(os.Stdin)
		if len(words) == 0 {
			fmt.Println("Too many or too few words!")
			return
		}
		if words[0] == "q" {
			return
		}

		if words[0] == "dict" || words[0] == "dictionary" {
			err = state.Dictionary(words[1])
			if err != nil {
				log.Fatalf("dictionary failed: %s", err)
			}
			continue
		}

		id, err := strconv.Atoi(words[0])
		if err != nil {
			fmt.Printf("This is not a number nor valid command!: %s", err)
			continue
		}

		if id > pageCount || id < 1 {
			fmt.Println("This page does not exist!")
			continue
		}

		pageDir := path + fmt.Sprintf("page%s/", words[0])
		pathFile := pageDir + "page.txt"

		imagePath := pageDir + "image.png"
		cmd := exec.Command("wslview", imagePath)
		if err := cmd.Start(); err == nil {
			fmt.Printf("imagePath %s\n", imagePath)
		}
		fmt.Printf("%s", err)

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
