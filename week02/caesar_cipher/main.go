package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Print("text: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text := strings.ToUpper(scanner.Text())
	fmt.Print("key: ")
	var key int
	fmt.Scan(&key)
	fmt.Println(caesarCipher(text, key))
}

func caesarCipher(text string, key int) string {
	key %= 26
	if key < 0 {
		key += 26
	}
	keyRune := rune(key)
	text = strings.ToUpper(text)
	var sb strings.Builder
	for _, r := range text {
		if 'A' <= r && r <= 'Z' {
			sb.WriteRune('A' + ((r - 'A' + keyRune) % 26))
		} else {
			sb.WriteRune(r)
		}
	}
	return sb.String()
}
