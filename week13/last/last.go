package main

import "fmt"

func Last[T any](values []T) T {
	return values[len(values)-1]
}

func main() {
	numbers := []int{10, 20, 30}
	words := []string{"go", "is", "fun"}

	fmt.Println(Last(numbers)) // should be "30"
	fmt.Println(Last(words))   // should be "fun"
}
