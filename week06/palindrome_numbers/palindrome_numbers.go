package palindromenumbers

import (
	"strconv"
)

func IsPalindrome(n int) bool {
	// ignore negative sign
	if n < 0 {
		n = -n
	}
	p1 := IsPalindromeString(strconv.Itoa(n))
	p2 := IsPalindromeSlice(slicedInt(n))
	p3 := IsPalindromeInt(n)
	return p1 && p2 && p3
}

func IsPalindromeString(s string) bool {
	r := []rune(s)
	for i, j := 0, len(r)-1; i <= j; i, j = i+1, j-1 {
		if r[i] != r[j] {
			return false
		}
	}
	return true
}

func IsPalindromeSlice(s []int) bool {
	for i, j := 0, len(s)-1; i <= j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}

func IsPalindromeInt(n int) bool {
	return n == reversedInt(n)
}

// slicedInt slices an integer into a slice of digits in reverse order
func slicedInt(n int) []int {
	s := []int{}
	for n > 0 {
		d := n % 10
		s = append(s, d)
		n /= 10
	}
	return s
}

// reversedInt reverses an integer
func reversedInt(n int) int {
	b := 0
	for n > 0 {
		d := n % 10
		n /= 10
		b *= 10
		b += d
	}
	return b
}
