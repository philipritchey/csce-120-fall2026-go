package main

import "fmt"

// enter a character: A
// you entered: 'A'
// unicode code point as decimal: 65
// unicode code point as hex: U+0041
// type: int32
// fits in a byte: true

func main() {
	fmt.Print("enter a character: ")
	var char rune
	fmt.Scanf("%c", &char)
	fmt.Printf("you entered: %q\n", char)
	fmt.Printf("unicode code point as decimal: %d\n", char)
	fmt.Printf("unicode code point as hex: %U\n", char)
	fmt.Printf("type: %T\n", char)
	fmt.Printf("fits in a byte: %t\n", 0 <= char && char <= 255)
}
