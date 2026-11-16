package main

import "fmt"

func main() {
	fmt.Println(Factorial(5))
}

func Factorial(n int) int {
	fmt.Printf("Factorial(%d) = ", n)
	if n < 0 {
		fmt.Println("0")
		return 0
	}
	if n == 0 {
		fmt.Println("1")
		return 1
	}
	fmt.Printf("%d * Factorial(%d)\n", n, n-1)
	return n * Factorial(n-1)
}
