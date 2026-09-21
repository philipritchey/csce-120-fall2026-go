package increaseaverage

import (
	"fmt"
	"io"
	"os"
)

func main() {
	Main(os.Stdin, os.Stdout)
}

func Main(r io.Reader, w io.Writer) {
	var ntamu, ntu int
	fmt.Fscan(r, &ntamu, &ntu)
	a := readInts(r, ntamu)
	b := readInts(r, ntu)
	fmt.Fprintln(w, averageIncreaseCount(a, b))
}

// readInts reads n integers from r and returns them in a slice
func readInts(r io.Reader, n int) []int {
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(r, &a[i])
	}
	return a
}

// averageIncreaseCount counts how many elements of a could move to b
// and increase the average of both.
func averageIncreaseCount(a, b []int) int {
	avgA := average(a)
	avgB := average(b)
	cnt := 0
	for _, v := range a {
		vF := float64(v)
		if avgB < vF && vF < avgA {
			cnt++
		}
	}
	return cnt
}

// average computes the average of the values in the slice
func average(a []int) float64 {
	return float64(sum(a)) / float64(len(a))
}

// sum computes the sum of the values in the slice
func sum(a []int) int {
	s := 0
	for _, v := range a {
		s += v
	}
	return s
}
