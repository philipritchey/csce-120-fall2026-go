package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	c := 0

	for range 1000 {
		wg.Go(func() {
			c = c + 1
		})
	}

	wg.Wait()
	fmt.Println(c)
}
