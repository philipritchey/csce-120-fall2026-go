package sumsquaredifference

func SumOfSquares(n int) int {
	return (n * (n + 1) * (2*n + 1)) / 6
}

func SquareOfSum(n int) int {
	s := (n * (n + 1)) / 2
	return s * s
}

func SumSquareDifference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
