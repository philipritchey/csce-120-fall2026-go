package main

import "fmt"

func Middle[T any](values []T) (T, bool) {
	if len(values) == 0 {
		var zero T
		return zero, false
	}
	return values[len(values)/2], true
}

func main() {
	fmt.Println(Middle([]int{10, 20, 30}))            // should be "20"
	fmt.Println(Middle([]string{"a", "b", "c", "d"})) // should be "c"
	fmt.Println(Middle([]float64{}))                  // should not panic
}
