package main

import "fmt"

func makeNumber() *int {
	x := 42
	return &x
}

func main() {
	p := makeNumber()
	fmt.Println(*p)
}
