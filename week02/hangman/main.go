package main

import (
	"fmt"
	"strings"
)

func main() {
	// read word and guesses
	var word, guesses string
	fmt.Scan(&word, &guesses)

	// for each guess, count number of hits, or a miss
	// win if hits == len(word)
	// lose if misses == 10
	var win bool
	var hits, misses int
	for _, r := range guesses {
		if strings.ContainsRune(word, r) {
			hits += strings.Count(word, string(r))
			if hits == len(word) {
				win = true
				break
			}
		} else {
			misses++
			if misses == 10 {
				win = false
				break
			}
		}
	}
	if win {
		fmt.Println("WIN")
	} else {
		fmt.Println("LOSE")
	}
}
