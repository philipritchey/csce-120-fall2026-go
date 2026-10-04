package main

import "fmt"

func main() {
	tests := []struct {
		a string
		b string
		i int
	}{
		{"apple", "banana", 0},
		{"cat", "car", 2},
		{"pre", "prefix", 3},
		{"howdy", "howdy", -1},
		{"😀😀😁", "😀😀😅", 2},
	}
	for _, t := range tests {
		j := FirstMismatch(t.a, t.b)
		if j != t.i {
			fmt.Printf("[FAIL] expected %q and %q to first differ at %d, got %d\n", t.a, t.b, t.i, j)
		}
	}
}

func FirstMismatch(a, b string) int {
	ra := []rune(a)
	rb := []rune(b)
	i := 0
	for i < len(ra) && i < len(rb) && ra[i] == rb[i] {
		i++
	}
	if i < len(ra) {
		return i
	} else if i < len(rb) {
		return i
	}
	return -1
}
