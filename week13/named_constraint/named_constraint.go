package main

import "fmt"

type Ordered interface {
	int | float64 | string
}

// func Max[T int | float64 | string](a, b T) T {
func Max[T Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

// func Min[T int | float64 | string](a, b T) T {
func Min[T Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(Max(10, 20))       // should be "20"
	fmt.Println(Min(2.1, 3.5))     // should be "2.1"
	fmt.Println(Max("cat", "dog")) // should be "dog"
}
