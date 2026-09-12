package main

import "fmt"

func main() {
	n := readInt()
	var v1, v2 int
	for range n {
		v := readInt()
		v1, v2 = topTwo(v1, v2, v)
	}
	fmt.Println(v2)
}

func readInt() int {
	var n int
	fmt.Scan(&n)
	return n
}

func topTwo(v1, v2, v3 int) (int, int) {
	switch {
	case v3 > v1:
		return v3, v1
	case v3 > v2:
		return v1, v3
	}
	return v1, v2
}
