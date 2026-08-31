package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

// enter your name: prof. coolguy ritchey
// your initials are P.C.R.

func main() {
	// read a line of text
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		line := scanner.Text()
		// uppercase it
		line = strings.ToUpper(line)
		// split on spaces
		names := strings.Split(line, " ")
		// for each word/name
		for _, name := range names {
			// get first character
			initial, _ := utf8.DecodeRuneInString(name)
			// print followed by .
			fmt.Printf("%c.", initial)
		}
		// print newline
		fmt.Println()
	}
}
