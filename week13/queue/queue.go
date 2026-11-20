package main

import "fmt"

type Queue[T any] struct {
	items []T
}

func (q *Queue[T]) Enqueue(item T) {
	q.items = append(q.items, item)
}

func (q *Queue[T]) Dequeue() (T, bool) {
	if q.IsEmpty() {
		var zero T
		return zero, false
	}

	item := q.items[0]
	q.items = q.items[1:]
	return item, true
}

func (q *Queue[T]) IsEmpty() bool {
	return len(q.items) == 0
}

func main() {
	demoQueueOfInt()
	fmt.Println()
	demoQueueOfString()
}

func demoQueueOfInt() {
	var q Queue[int]
	q.Enqueue(8)
	count := 0
	for {
		value, ok := q.Dequeue()
		if !ok {
			// queue is empty
			fmt.Println("failure.")
			break
		}
		count++
		if value == 9 {
			fmt.Println("success!")
			break
		}
		if value%6 == 4 {
			q.Enqueue((value - 1) / 3)
		}
		q.Enqueue(value * 2)
	}
	fmt.Printf("explored %d values\n", count)
}

func demoQueueOfString() {
	var q Queue[string]
	v := make(map[string]bool)
	q.Enqueue("cat")
	count := 0
	for {
		value, ok := q.Dequeue()
		if !ok {
			// queue is empty
			fmt.Println("failure.")
			break
		}
		if v[value] {
			// already visited
			continue
		}
		count++
		if value == "dog" {
			fmt.Println("success!")
			break
		}
		letters := []byte(value)
		for i := range letters {
			if letters[i] > 'a' {
				letters[i]--
				q.Enqueue(string(letters))
				letters[i]++
			}
			if letters[i] < 'z' {
				letters[i]++
				q.Enqueue(string(letters))
				letters[i]--
			}
		}
		v[value] = true
	}
	fmt.Printf("explored %d values\n", count)
}
