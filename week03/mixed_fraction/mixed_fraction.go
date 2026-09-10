package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	for a > 0 || b > 0 {
		fmt.Println(asMixedFraction(a, b))
		fmt.Scan(&a, &b)
	}
}

func asMixedFraction(a, b int) string {
	q := a / b
	r := a % b
	r, b = reduceFraction(r, b)
	var s string
	if q > 0 {
		s += fmt.Sprintf("%d ", q)
	}
	if r > 0 {
		s += fmt.Sprintf("%d / %d", r, b)
	}
	return s
}

func reduceFraction(a, b int) (int, int) {
	g := gcd(a, b)
	return a / g, b / g
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
