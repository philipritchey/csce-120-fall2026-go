package main

import (
	"fmt"
)

// Write a program that reads a single character and identifies whether it is:
// * lowercase letter
// * uppercase letter
// * digit
// * something else

func main() {
	// read a character
	var c rune
	fmt.Scanf("%c", &c)
	fmt.Printf("%q is ", c)
	switch {
	case isLowercase(c):
		fmt.Println("a lowercase letter")
	case isUppercase(c):
		fmt.Println("an uppercase letter")
	case isDigit(c):
		fmt.Println("a digit")
	default:
		fmt.Println("something else")
	}
}

func isLowercase(c rune) bool {
	return inRuneRange(c, 'a', 'z')
}

func isUppercase(c rune) bool {
	return inRuneRange(c, 'A', 'Z')
}

func isDigit(c rune) bool {
	return inRuneRange(c, '0', '9')
}

func inRuneRange(r, lower, upper rune) bool {
	return lower <= r && r <= upper
}
