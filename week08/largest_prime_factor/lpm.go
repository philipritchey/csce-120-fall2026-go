package main

import "fmt"

func main() {
	actual := lpf(13195)
	if actual != 29 {
		fmt.Println("expected LPF(13195) to be 29, got", actual)
		return
	}

	actual = lpf(4567 * 8191)
	if actual != 8191 {
		fmt.Println("expected LPF(", 4567*8191, ") to be 8191, got", actual)
		return
	}

	fmt.Println("LPF(600851475143) =", lpf(600851475143))
}

// the largest prime factor of n
func lpf(n int) int {
	p := 2
	for p < n {
		for n%p == 0 {
			n /= p
		}
		p += 1
	}
	return p // == n
}
