package main

import "fmt"

type Pair[T1, T2 any] struct {
	first  T1
	second T2
}

func main() {
	p1 := Pair[int, int]{1, 2}
	fmt.Println(p1)

	p2 := Pair[string, bool]{"howdy", true}
	fmt.Println(p2)

	fmt.Printf(" first: %v\nsecond: %v\n", p2.first, p2.second)
}
