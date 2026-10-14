package main

import "fmt"

func main() {
	a := 5
	b := 8

	p := &a
	q := &b

	*p = *p + 1
	*q = *p + *q

	fmt.Println(a)
	fmt.Println(b)
}
