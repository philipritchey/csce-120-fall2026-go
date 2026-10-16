package main

import (
	"fmt"
)

func main() {
	n, p, q := lpp()
	fmt.Println(n, "=", p, "*", q)
	if n != 906609 {
		fmt.Println("expected 906609 = 993 * 913")
	}
}

func lpp() (int, int, int) {
	n2, p2, q2 := 0, 0, 0
	p := 999
	for p > 0 && p*999 > n2 {
		q := 999
		for q >= p {
			n := p * q
			if isPalindrome(n) && n > n2 {
				n2, p2, q2 = n, p, q
			}
			q--
		}
		p--
	}
	return n2, p2, q2
}

func isPalindrome(n int) bool {
	return n == reversed(n)
}

func reversed(n int) int {
	r := 0
	for n > 0 {
		d := n % 10
		n /= 10
		r *= 10
		r += d
	}
	return r
}
