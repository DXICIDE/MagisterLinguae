package repl

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// scans the input for multiple lines
func Scan() []string {
	fmt.Print("> ")
	scanner := bufio.NewScanner(os.Stdin)
	scanned := scanner.Scan()
	if !scanned {
		return nil
	}
	line := scanner.Text()
	line = strings.TrimSpace(line)
	return strings.Fields(line)
}
