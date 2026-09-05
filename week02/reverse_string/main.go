package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	s := scanner.Text()
	r := reverse(s)
	fmt.Println(r)
}

func reverse(s string) string {
	var r string
	for _, c := range s {
		r = string(c) + r
	}
	return r
}
