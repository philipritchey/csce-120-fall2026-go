package main

import "fmt"

func main() {
	var x [4]int
	fmt.Scan(&x[0], &x[1], &x[2], &x[3])
	fmt.Println(CombinationLock(x[0], x[1], x[2], x[3]))
}
