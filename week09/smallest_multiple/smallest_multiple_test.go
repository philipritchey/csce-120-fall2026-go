package smallestmultiple

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/require"
)

// 2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.
// What is the smallest positive number that is evenly divisible (divisible with no remainder) by all of the numbers from 1 to 20?

// SmallestMultiple returns the smallest positive number that is evenly divisible
// by all of the numbers from 1 to n
func SmallestMultiple(n int) int {
	// get list of primes up to n
	primes := sieve(n)
	// for each prime p_i, take k_i, where p_i^k_i <= n
	product := 1
	for _, p := range primes {
		m := n
		for m >= p {
			product *= p
			m /= p
		}
	}
	// return product of p_i^k_i
	return product
}

// sieve returns a list of the primes up to n
// e.g. sieve(10) -> [2 3 5 7]
func sieve(n int) []int {
	s := slices.Repeat([]bool{true}, n+1)
	s[0], s[1] = false, false
	for i := 2; i*i <= n; i++ {
		if s[i] {
			for j := i * i; j <= n; j += i {
				s[j] = false
			}
		}
	}
	primes := make([]int, 0)
	for p, b := range s {
		if b {
			primes = append(primes, p)
		}
	}
	return primes
}

// the smallest positive number that is evenly divisible
// by all of the numbers from 1 to n
func SmallestMultiple2(n int) int {
	if n <= 2 {
		return n
	}
	return lcm(n, SmallestMultiple2(n-1))
}

// least common multiple
func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

// greatest common divisor
func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func TestSieve(t *testing.T) {
	require.ElementsMatch(t, []int{2, 3, 5, 7}, sieve(10))
}

func TestSmallestMultiple(t *testing.T) {
	// 1 * 2 * 3 * 4 * 5 * 6 * 7 * 8 * 9 * 10 = 3628800
	// primes: 2{3},3{2},5{1},7{1}
	//          8    9    5    7
	// 8 * 9 * 5 * 7 = 2520
	require.Equal(t, 2_520, SmallestMultiple(10), "n = 10")

	// 1 * 2 * 3 * 4 * 5 * 6 * 7 * 8 * 9 * 10
	// 11 * 12 * 13 * 14 * 15 * 16 * 17 * 18 * 19 * 20
	// primes: 2{4},3{2},5{1},7{1},11{1},13{1},17{1},19{1}
	//          16   9    5    7    11     13   17   19
	// 16 * 9 * 5 * 7 * 11 * 13 * 17 * 19 = 232792560
	require.Equal(t, 232_792_560, SmallestMultiple(20), "n = 20")

	// an alternative method
	// lcm(1,2,...,10) = 2520
	// lcm(1,2,...,n) = lcm(lcm(1,2,...,n-1), n)
	// lcm(a, b) = a*b / gcd(a, b)
	require.Equal(t, 2_520, SmallestMultiple2(10), "(2) n = 10")
	require.Equal(t, 232_792_560, SmallestMultiple2(20), "(2) n = 20")
}
