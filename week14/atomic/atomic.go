package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var wg sync.WaitGroup
	var c int32 = 0 // updated

	for range 1000 {
		wg.Go(func() {
			atomic.AddInt32(&c, 1) // updated
		})
	}

	wg.Wait()
	fmt.Println(c)
}
