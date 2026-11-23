package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	c := 0

	var mu sync.Mutex // added

	for range 1000 {
		wg.Go(func() {
			mu.Lock()   // added
			c = c + 1   // only one goroutine at a time here
			mu.Unlock() // added
		})
	}

	wg.Wait()
	fmt.Println(c)
}
