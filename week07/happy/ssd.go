package happy

// Ssd computes the Sum of Squared Digits of n in base b.
// Must have: b > 1 and n >= 0
// Example: Ssd(10, 1234) = 1^2 + 2^2 + 3^2 + 4^2 = 30
func Ssd(b, n int) int {
	s := 0
	for n > 0 {
		d := n % b
		s += d * d
		n /= b
	}
	return s
}
