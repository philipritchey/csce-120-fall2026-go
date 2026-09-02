package main

import "fmt"

// Write a program that validates a course code
// * A course code is department code and course number
//   * A valid department code is an uppercase letter
//   * A valid course number is between 100 and 499

func main() {
	// read a course code (char, integer)
	var dept rune
	var num int
	fmt.Scanf("%c%d", &dept, &num)
	// validate it
	fmt.Printf("valid: %t\n", isValid(dept, num))
}

// check dept code is uppercase lettter and course number in range [100, 499]
func isValid(dept rune, num int) bool {
	return isUppercase(dept) && 100 <= num && num <= 499
}

func isUppercase(c rune) bool {
	return 'A' <= c && c <= 'Z'
}
