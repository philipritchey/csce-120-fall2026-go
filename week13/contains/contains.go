package main

import (
	"fmt"
)

// func ContainsInt(values []int, target int) bool {
// 	for _, value := range values {
// 		if value == target {
// 			return true
// 		}
// 	}
// 	return false
// }

// func ContainsString(values []string, target string) bool {
// 	for _, value := range values {
// 		if value == target {
// 			return true
// 		}
// 	}
// 	return false
// }

func Contains[T comparable](values []T, target T) bool {
	for _, value := range values {
		if value == target {
			return true
		}
	}
	return false
}

// ooo... ahhh.... a generic test function =)
func ExpectContains[T comparable](values []T, target T) {
	if Contains(values, target) {
		fmt.Printf("[ OK ] Contains(%v, %v) is true\n", values, target)
	} else {
		fmt.Printf("[FAIL] Contains(%v, %v) should be true, got false\n", values, target)
	}
}

func ExpectNotContains[T comparable](values []T, target T) {
	if Contains(values, target) {
		fmt.Printf("[FAIL] Contains(%v, %v) should be false, got true\n", values, target)
	} else {
		fmt.Printf("[ OK ] Contains(%v, %v) is false\n", values, target)
	}
}

func main() {
	ExpectNotContains([]int{1, 2, 3}, 4)
	ExpectContains([]int{1, 2, 3}, 2)
	ExpectContains([]string{"one", "two", "three"}, "two")
	ExpectNotContains([]string{"one", "two", "three"}, "four")

	ExpectContains([]int{1, 2, 3}, 2)
	ExpectNotContains([]string{"red", "green", "blue"}, "yellow")
	ExpectContains([]bool{true, false}, true)

	// doesn't work becase []int not comparable
	// ExpectContainsTrue([][]int{{1, 2}, {3, 4}}, []int{1, 2})

}
