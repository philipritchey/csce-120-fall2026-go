package main

import (
	"fmt"
	"math/rand/v2"
	"slices"
)

func main() {
	items := randomItems(100000)

	c := make(chan int) // Allocate a channel.
	// Start the sort in a goroutine; when it completes, signal on the channel.
	fmt.Println("main kicks off the goroutine")
	go func() {
		fmt.Println("goroutine start")
		slices.Sort(items)
		fmt.Println("goroutine end")
		c <- 1 // Send a signal; value does not matter.
	}()
	doSomethingElseForAWhile()
	fmt.Println("main is waiting...")
	<-c // Wait for sort to finish; discard sent value.
	doSomethingThatUsesTheSortedList(items)
}

func randomItems(n int) []int {
	items := make([]int, n)
	for i := range n {
		items[i] = rand.Int()
	}
	return items
}

func doSomethingElseForAWhile() {
	fmt.Println("doing something else for a while")
}

func doSomethingThatUsesTheSortedList(items []int) {
	fmt.Println("doing something that uses the sorted list")
	i, j := 0, len(items)-1
	target := items[j] + 1
	s := items[i] + items[j]
	for i < j && s != target {
		if s < target {
			i++
		} else {
			j--
		}
		s = items[i] + items[j]
	}
	fmt.Println("success:", i < j)
}
