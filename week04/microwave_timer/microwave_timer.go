package main

import "fmt"

func main() {
	var n int
	_, err := fmt.Scan(&n)
	if err != nil || n < 0 || n > 9999 {
		fmt.Println("invalid input, expected integer in range [0, 9999]")
		return
	}
	fmt.Println(MicrowaveTimer(n))
}

func MicrowaveTimer(t int) int {
	// note: this was super-easy to code AFTER writing the test cases
	s := t % 100
	m := t / 100
	return 60*m + s
}
