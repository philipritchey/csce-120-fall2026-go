package main

import "fmt"

func Add[T int | float64](a, b T) T {
	return a + b
}

func main() {
	fmt.Println(Add(3, 4))     // should be "7"
	fmt.Println(Add(2.3, 1.6)) // should be "3.9"
	// fmt.Println(Add(true, false)) // invalid
}
