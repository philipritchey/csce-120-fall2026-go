package main

import "fmt"

func main() {
	// read n, the number of numbers to expect
	var n int
	fmt.Scan(&n)
	// init sum to 0
	sum := 0.0
	// do n times
	for range n {
		// read a number
		var value float64
		fmt.Scan(&value)
		// add value to sum
		sum += value
	}
	// divide sum by n to get average
	average := sum / float64(n)
	// output average with 4 digits of precision
	fmt.Printf("%.4f\n", average)
}
