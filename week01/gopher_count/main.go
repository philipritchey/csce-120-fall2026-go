package main

import "fmt"

func main() {
	// modify this example from [Go by Example: For](https://gobyexample.com/for)
	// for i := range 3 {
	//     fmt.Println("range", i)
	// }
	for i := range 5 { // change 3 -> 5
		fmt.Println(i+1, "Gopher") // swap order, change i->i+1, change "range"->"Gopher"
	}
}
