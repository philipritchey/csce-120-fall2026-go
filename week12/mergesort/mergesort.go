package main

import (
	"fmt"
	"math/rand/v2"
	"strings"
)

func main() {
	n := 15
	slice := make([]int, n)
	for i := range n {
		slice[i] = rand.IntN(100)
	}
	fmt.Println(slice)
	fmt.Println(MergeSort(slice, 0))
}

func MergeSort(slice []int, depth int) []int {
	fmt.Print(strings.Repeat("| ", depth))
	fmt.Print("merge sort ", slice, ": ")
	if len(slice) <= 1 {
		fmt.Println("done")
		return slice
	}
	mid := len(slice) / 2
	left := slice[:mid]
	right := slice[mid:]
	fmt.Println(left, right)
	return Merge(MergeSort(left, depth+1), MergeSort(right, depth+1), depth+1)
}

// merge the two sorted lists into one sorted list
func Merge(a, b []int, depth int) []int {
	fmt.Print(strings.Repeat("| ", depth))
	fmt.Print("merge ", a, " ", b, ": ")
	c := []int{}
	i, j := 0, 0
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			c = append(c, a[i])
			i++
		} else {
			c = append(c, b[j])
			j++
		}
	}
	c = append(c, a[i:]...)
	c = append(c, b[j:]...)
	fmt.Println(c)
	return c
}
