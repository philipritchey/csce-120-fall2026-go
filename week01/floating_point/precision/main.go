package main

import "fmt"

func main() {
	var x float32 = 0.1
	var y float64 = 0.1
	var z float64 = float64(x)
	fmt.Printf("x: %.8f%8s = %.27f\n", x, "", x)
	fmt.Printf("y: %.16f = %.55f\n", y, y)
	fmt.Printf("z: %.16f = %.55f\n", z, z)
}
