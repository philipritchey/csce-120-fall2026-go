package main

func ReverseWithStack[T any](values []T) []T {
	var stack Stack[T]

	// push all values
	for _, value := range values {
		stack.Push(value)
	}

	// pop all values into result
	result := make([]T, len(values))
	for i := range result {
		value, ok := stack.Pop()
		if !ok {
			panic("unexpextedly empty stack")
		}
		result[i] = value
	}

	return result
}
