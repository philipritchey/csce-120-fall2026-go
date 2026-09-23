package main

// CombinationLock determines how many degrees a combination lock dial is rotated in total
// to open the lock
func CombinationLock(initial, first, second, third int) int {
	const numbers = 40 // on the dial
	const degreesPerNumber = 360 / numbers
	return degreesPerNumber * (3*numbers + mod(initial-first, numbers) + mod(second-first, numbers) + mod(second-third, numbers))
}

// mod computes the non-negative value of a mod b
// b must be positive
func mod(a, b int) int {
	m := a % b
	if m < 0 {
		m += b
	}
	return m
}
