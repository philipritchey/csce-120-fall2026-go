package main

import "fmt"

func double(n *int) {
	*n = *n * 2
}

func main() {
	x := 7

	double(&x)

	fmt.Println(x)
}
