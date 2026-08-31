package main

import "fmt"

// enter an integer: 8675
// ascii: invalid
// unicode: U+21E3 '⇣'

func main() {
	fmt.Print("enter an integer: ")
	var i int32
	_, err := fmt.Scan(&i)
	if err != nil {
		panic(err)
	}
	fmt.Print("ascii: ")
	if 0 <= i && i <= 255 {
		fmt.Printf("%c\n", i)
	} else {
		fmt.Print("invalid\n")
	}
	fmt.Print("unicode: ")
	if i >= 0 {
		fmt.Printf("%#U\n", i)
	} else {
		fmt.Print("invalid\n")
	}
}
