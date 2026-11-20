package main

import "fmt"

func main() {
	var numbers Stack[int]

	numbers.Push(10)
	numbers.Push(20)
	numbers.Push(30)

	value, ok := numbers.Pop()
	if ok {
		fmt.Println(value)
	}

	var words Stack[string]

	words.Push("first")
	words.Push("second")

	word, ok := words.Pop()
	if ok {
		fmt.Println(word)
	}

	fmt.Println(ReverseWithStack([]int{1, 2, 3, 4}))       // should be "[4 3 2 1]"
	fmt.Println(ReverseWithStack([]string{"a", "b", "c"})) // should be "[c b a]"
	fmt.Println(ReverseWithStack([]bool{}))                // should be "[]"
}
