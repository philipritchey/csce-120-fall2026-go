package main

import "fmt"

func main() {
	slice := []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377}
	fmt.Println(BinarySearch(slice, 1))
	fmt.Println(BinarySearch(slice, 13))
	fmt.Println(BinarySearch(slice, 22))
}

func BinarySearch(slice []int, target int) bool {
	fmt.Print("searching ", slice, " for ", target, ": ")
	if len(slice) <= 1 {
		return len(slice) > 0 && slice[0] == target
	}
	mid := len(slice) / 2
	left := slice[:mid]
	right := slice[mid:]
	if target < slice[mid] {
		fmt.Println(target, "is in left half")
		return BinarySearch(left, target)
	} else {
		fmt.Println(target, "is in right half")
		return BinarySearch(right, target)
	}
}
