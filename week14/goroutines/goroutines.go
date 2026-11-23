package main

import (
	"fmt"
	"time"
)

func doSomething(i int) {
	fmt.Println(i, "is doing something...")
}

func main() {
	go doSomething(1)
	go doSomething(2)
	go doSomething(3)
	time.Sleep(time.Millisecond)
	fmt.Println("main is done.")
}
