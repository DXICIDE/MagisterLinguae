package repl

import "fmt"

func Help() {
	fmt.Println("usage: [command] <path> <text>")
	fmt.Println("\nHere are all commands and their use case")
	fmt.Println("learn <text>		Will output words based on if they are known or not")
	fmt.Println("learnfile <path>	Will output words based on if they are known or not")
	fmt.Println("mark <text>		Mark the word as known")
	fmt.Println("unmark <text>		Unmark the word as known")
	fmt.Println("lookup <word>		Show stats of the word")
	fmt.Println("switch			Switch language")
	fmt.Println("list			Show list of the word by frequency and known status")
	fmt.Println("resetdb 		Reset the database!!!")
	fmt.Println("help 			print help")
	fmt.Println("newlang 		prompts you to create new language tab with custom code and name")
	fmt.Println("lang 			prints the current used language tab")
	fmt.Println("dict / dictionary <word> 		lists the word and all it's available meaning")
	fmt.Println("deletelang 		delete language by typing code")
	fmt.Println("anki 			Gives u a quiz, similar to anki. Only on words u marked as known, with no translate")
	fmt.Println("q / quit 		quit the repl")
	fmt.Println("deephelp 		List all the command and their usage and purpose")
}

func DeepHelp() {
	fmt.Println("\nHere are all commands and their purpose")
	fmt.Println("learn <text>		Will output words based on if they are known or not")
	fmt.Println("learnfile <path>	Will output words based on if they are known or not")
	fmt.Println("mark <text>		Mark the word as known, reccomend that you mark only the words you 100% know the meaning of")
	fmt.Println("unmark <text>		Unmark the word as known, for words you have forgotten or realised you werent sure of their meaning as before")
	fmt.Println("lookup <word>		Show stats of the word, good for knowing when was the last time you seen that word and other stats like frequency or known status")
	fmt.Println("switch			Switch language to supported ones, the abylity to create your own will come later")
	fmt.Println("list			Show list of the word by frequency and known status, very important for searching for frequent words that are you don't know, should make learning the language much easier")
	fmt.Println("resetdb 		Reset the database!!! This is a complete wipe of db, no warning, you will lose everything, this will be changed later")
	fmt.Println("help 			Print help")
	fmt.Println("newlang 		prompts you to create new language tab with custom code and name")
	fmt.Println("lang 			prints the current used language tab")
	fmt.Println("dict / dictionary <word> 		lists the word and all it's available meaning, only available for single words, input theory discourages translationg the entire meaning of sentences, thats why i wont even be implementing this option")
	fmt.Println("deletelang 		delete language by typing code")
	fmt.Println("anki 			Gives u a quiz, similar to anki. Only on words u marked as known, with no translate. The purpose is to prompt you to answer if u know the words in isolation as if it was just taken out of sentence")
	fmt.Println("q / quit 		Quit the repl")
	fmt.Println("deephelp 		List all the command and their usage and purpose")
}
