package main

import "fmt"

type Counter struct {
	Value int
}

func (c *Counter) Increment() {
	c.Value++
}

func (c *Counter) Add(n int) {
	c.Value += n
}

func (c *Counter) Reset() {
	c.Value = 0
}

func (c Counter) String() string {
	return fmt.Sprintf("%d", c.Value)
}

func main() {
	count := Counter{Value: 10}

	fmt.Println("Starting value:", count)

	count.Increment()
	fmt.Println("After Increment:", count)

	count.Add(5)
	fmt.Println("After Add(5):", count)

	count.Reset()
	fmt.Println("After Reset:", count)
}
