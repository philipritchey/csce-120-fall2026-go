package main

import (
	"fmt"
	"strconv"
)

// Count from 1 to N, saying each number out loud except
// If a number is divisible by 2, say "Fizz" instead
// If a number is divisible by 3, say "Buzz" instead
// If a number is divisible by 5, say "Jazz" instead
// If a number is divisible by more than one of 2, 3, and 5,
//   say both/all of "Fizz", "Buzz", or "Jazz"

func main() {
	// read an integer n
	var n int
	fmt.Scan(&n)
	// count from i=1 to n
	for i := 1; i <= n; i++ {
		var s string
		//   if n is divisible by 2 -> "Fizz"
		if i%2 == 0 {
			s += "Fizz"
		}
		//   if n is divisible by 3 -> "Buzz"
		if i%3 == 0 {
			s += "Buzz"
		}
		//   if n is divisible by 5 -> "Jazz"
		if i%5 == 0 {
			s += "Jazz"
		}
		//   else (none of the above) -> i
		if len(s) == 0 {
			s = strconv.Itoa(i)
		}
		fmt.Println(s)
	}
}
