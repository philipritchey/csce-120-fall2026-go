package main

import "fmt"

func main() {
	x := 3

	p := &x
	q := p

	*q = 10

	fmt.Println(x)
	fmt.Println(*p)
	fmt.Println(*q)
}
