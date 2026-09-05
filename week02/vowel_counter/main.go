package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	s := scanner.Text()
	fmt.Println(vowelCount(s))
	fmt.Println(vowelCountAlt1(s))
	fmt.Println(vowelCountAlt2(s))
}

func vowelCount(s string) int {
	c := 0
	for _, r := range strings.ToLower(s) {
		switch r {
		case 'a', 'e', 'i', 'o', 'u':
			c++
		}
	}
	return c
}

func vowelCountAlt1(s string) int {
	s = strings.ToLower(s)
	return strings.Count(s, "a") +
		strings.Count(s, "e") +
		strings.Count(s, "i") +
		strings.Count(s, "o") +
		strings.Count(s, "u")
}

func vowelCountAlt2(s string) int {
	vowels := "aeiouAEIOU"
	c := 0
	for _, r := range s {
		if strings.ContainsRune(vowels, r) {
			c++
		}
	}
	return c
}
