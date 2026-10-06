package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	// read the content of quotes.txt
	bytes, err := os.ReadFile("quotes.txt")
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}

	text := string(bytes)
	runes := []rune(text)
	lineCount := strings.Count(text, "\n") + 1
	charCount := len(runes)
	summary := fmt.Sprintf("File summary\nLines: %d\nCharacters: %d", lineCount, charCount)

	err = os.WriteFile("summary.txt", []byte(summary), 0644)
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}
	fmt.Println("Wrote summary.txt")
}
