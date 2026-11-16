package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(Fibonacci(5, 0))
}

func Fibonacci(n, depth int) int {
	fmt.Print(strings.Repeat("| ", depth))
	fmt.Printf("Fibonacci(%d) = ", n)
	if n < 2 {
		fmt.Println(n)
		return n
	}
	fmt.Printf("Fibonacci(%d) + Fibonacci(%d)\n", n-1, n-2)
	return Fibonacci(n-1, depth+1) + Fibonacci(n-2, depth+1)
}
