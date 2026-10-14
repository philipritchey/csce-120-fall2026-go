package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Every variable is stored somewhere in memory
	temperature := 72
	fmt.Println("temperature:", temperature)

	pause()

	// A pointer stores a memory address
	var tempPtr *int

	// The & operator gets a variable's memory address
	tempPtr = &temperature
	fmt.Println("address:", tempPtr)

	pause()

	// The * operator follows a pointer to access the value at that address
	// Access could be a read
	fmt.Println("value through pointer:", *tempPtr)

	pause()

	// Access could be a write
	*tempPtr = 75

	fmt.Println("after changing through pointer:")
	fmt.Println("temperature:", temperature)
	fmt.Println("value through pointer:", *tempPtr)

	pause()

	// Two pointers can refer to the same variable
	anotherPtr := tempPtr
	*anotherPtr = 80

	fmt.Println("after changing through anotherPtr:")
	fmt.Println("temperature:", temperature)
	fmt.Println("value through tempPtr:", *tempPtr)
	fmt.Println("value through anotherPtr:", *anotherPtr)
}

func pause() {
	fmt.Printf("\n...paused...\n")
	s := bufio.NewScanner(os.Stdin)
	s.Scan()
}
