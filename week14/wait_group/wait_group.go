package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

func worker(id int) {
	fmt.Println("Worker", id, "started")
	time.Sleep(time.Duration(rand.IntN(1000)) * time.Millisecond)
	fmt.Println("Worker", id, "done")
}

func main() {
	var wg sync.WaitGroup

	for i := range 5 {
		wg.Go(func() {
			worker(i + 1)
		})
	}
	fmt.Println("Waiting...")
	wg.Wait()
	fmt.Println("All workers finished")
}
