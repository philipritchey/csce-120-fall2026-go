package main

import (
	"fmt"
)

func Equal[T comparable](a, b T) bool {
	return a == b
}

func main() {
	fmt.Println(Equal(3, 3))
	fmt.Println(Equal("go", "java"))
	fmt.Println(Equal(true, false))
	fmt.Println(Equal([2]int{1, 2}, [2]int{1, 2}))
}
