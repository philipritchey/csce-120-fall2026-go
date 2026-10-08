package happy

func IsHappy(n int) bool {
	// until n is 1 or 4
	for !(n == 1 || n == 4) {
		n = Ssd(10, n)
	}
	return n == 1
}
