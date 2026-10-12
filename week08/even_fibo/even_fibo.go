package main

import "fmt"

func main() {
	testcases := []struct {
		in  int
		out int
	}{
		{10, 10},    // 0+2+8 = 10
		{100, 44},   // 0+2+8+34 = 44
		{1000, 798}, // 0+2+8+34+144+610 = 798
	}
	for _, tc := range testcases {
		actual := SumOfEvenFibo(tc.in)
		expected := tc.out
		if actual != expected {
			fmt.Printf("[FAIL] expected SumOfEvenFibo(%d) to be %d, got %d\n", tc.in, expected, actual)
			return
		}
	}
	// if these test pass, then the next result is probably correct

	fmt.Println(SumOfEvenFibo(4e6))
}

func SumOfEvenFibo(limit int) int {
	a, b, sum := 0, 1, 0
	for b <= limit {
		if b%2 == 0 {
			sum += b
		}
		a, b = b, a+b
	}
	return sum
}
