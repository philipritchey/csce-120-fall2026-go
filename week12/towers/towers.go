package main

import "fmt"

func main() {
	RecursiveTowersOfHanoi(3, "A", "C", "B")
	fmt.Println("=======================")
	IterativeTowersOfHanoi(3, "A", "C", "B")
}

func RecursiveTowersOfHanoi(n int, src, dest, aux string) {
	if n > 0 {
		RecursiveTowersOfHanoi(n-1, src, aux, dest)
		fmt.Printf("move disk %d from %s to %s\n", n, src, dest)
		RecursiveTowersOfHanoi(n-1, aux, dest, src)
	}
}

func IterativeTowersOfHanoi(n int, src, dest, aux string) {
	type state struct {
		n, d           int
		src, dest, aux string
	}
	stack := []state{{n, n, src, dest, aux}}
	for len(stack) > 0 {
		s := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		n, d, src, dest, aux := s.n, s.d, s.src, s.dest, s.aux
		if n == 1 {
			fmt.Printf("move disk %d from %s to %s\n", d, src, dest)
		} else {
			stack = append(stack, state{n - 1, d - 1, aux, dest, src})
			stack = append(stack, state{1, d, src, dest, aux})
			stack = append(stack, state{n - 1, d - 1, src, aux, dest})
		}
	}
}
