package main

import "fmt"

func main() {
	fmt.Println(NthPrime(1))
	fmt.Println(NthPrime(2))
	fmt.Println(NthPrime(3))
	fmt.Println(NthPrime(4))
	fmt.Println(NthPrime(5))
	fmt.Println(NthPrime(10001))
}

// NthPrime returns the n-th prime
func NthPrime(n int) int {
	primes := []int{2, 3}
	// all primes > 3 are 6k +/- 1
	p := 5 // 6k - 1
	for n > len(primes) {
		if IsPrime(p, primes) {
			primes = append(primes, p)
		}
		if IsPrime(p+2, primes) {
			primes = append(primes, p+2)
		}
		p += 6
	}
	return primes[n-1]
}

// IsPrime returns true if n is prime
func IsPrime(n int, primes []int) bool {
	if n < 2 {
		return false
	}
	if n < 4 {
		return true
	}
	// check with previously discovered primes
	for _, p := range primes {
		if p*p > n {
			break
		}
		if n%p == 0 {
			return false
		}
	}
	return true
}
