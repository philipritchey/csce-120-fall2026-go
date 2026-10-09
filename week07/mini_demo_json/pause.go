package main

import (
	"bufio"
	"fmt"
	"os"
)

// Pause waits for user to press Enter key then returns
func Pause() {
	fmt.Printf("\n...paused...\n")
	s := bufio.NewScanner(os.Stdin)
	s.Scan()
}
