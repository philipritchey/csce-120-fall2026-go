package main

import "fmt"

func main() {
	// accumulated errors
	var x float32
	var y float64
	for range 10 {
		x += 0.1
		y += 0.1
	}
	fmt.Println(x, y)

	// comparison failure
	var a, b float64 = 0.1, 0.2
	fmt.Printf("0.1 + 0.2 == 0.3: %t\n", a+b == 0.3)
	// but it "works" with float32
	var c, d float32 = 0.1, 0.2
	fmt.Printf("0.1 + 0.2 == 0.3: %t\n", c+d == 0.3)

	// precision loss
	x = 17000000
	fmt.Printf("%.1f + 1.0 = %.1f\n", x, x+1.0)
}
