package main

import "fmt"

func First[T any](values []T) (T, bool) {
	if len(values) == 0 {
		var zero T
		return zero, false
	}
	return values[0], true
}

func main() {
	fmt.Println(First([]int{4, 5, 6}))                   // should be "4"
	fmt.Println(First([]string{"red", "green", "blue"})) // should be "red"
	fmt.Println(First([]float64{1.5, 2.5, 3.5}))         // should be "1.5"
	fmt.Println(First([]complex64{}))                    // should not panic
}
