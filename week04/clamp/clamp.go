package main

import (
	"fmt"
	"io"
	"os"
)

type Range struct {
	lower int
	upper int
}

func main() {
	r := readRange(os.Stdin)
	n := readInt(os.Stdin)
	for range n {
		i := readInt(os.Stdin)
		fmt.Println(clamp(r, i))
	}
}

func readRange(reader io.Reader) Range {
	var r Range
	fmt.Fscan(reader, &r.lower, &r.upper)
	return r
}

func readInt(reader io.Reader) int {
	var n int
	fmt.Fscan(reader, &n)
	return n
}

func clamp(r Range, i int) int {
	if i < r.lower {
		return r.lower
	}
	if i > r.upper {
		return r.upper
	}
	return i
}
