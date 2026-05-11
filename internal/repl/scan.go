package repl

import (
	"bufio"
	"fmt"
	"io"
	"strings"
	"text/scanner"
	"unicode"
)

// stdin scan
func (state *AppState) Scan(r io.Reader) []string {
	fmt.Print("> ")
	scan := bufio.NewScanner(r)
	if !scan.Scan() {
		return nil
	}
	return TokenizeString(scan.Text())
}

func TokenizeString(line string) []string {
	line = strings.TrimSpace(line)
	if line == "" {
		return nil
	}

	var s scanner.Scanner
	s.Init(strings.NewReader(line))
	s.Mode = scanner.ScanIdents | scanner.ScanInts | scanner.ScanFloats

	// Customize what characters are considered part of identifiers
	s.IsIdentRune = func(ch rune, i int) bool {
		return unicode.IsLetter(ch) || unicode.IsDigit(ch) || ch == '\''
	}

	var tokens []string
	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
		tokens = append(tokens, s.TokenText())
	}

	if len(tokens) == 0 {
		fmt.Println("Nothing was inputed")
	}

	return tokens
}
