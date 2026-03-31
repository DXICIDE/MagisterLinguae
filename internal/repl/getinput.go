package repl

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"text/scanner"
)

// scans the input for multiple lines
func Scan() []string {
	fmt.Print("> ")
	in := bufio.NewScanner(os.Stdin)
	scanned := in.Scan()
	if !scanned {
		return nil
	}
	line := in.Text()
	line = strings.TrimSpace(line)
	var s scanner.Scanner
	s.Init(strings.NewReader(line))
	s.Mode = scanner.ScanIdents | scanner.ScanStrings

	var tokens []string
	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
		tokens = append(tokens, s.TokenText())
	}
	return tokens
}
