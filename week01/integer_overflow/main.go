package main

import "fmt"

func main() {
	var a, b int8 = 123, 45
	// 123 + 45 "should be" 168
	// but 168 > 127 (max value for int8)
	fmt.Println(a, "+", b, "=", a+b)

	var c, d uint8 = 12, 34
	// 12 - 34 "should be" -22
	// but -22 < 0 (min value for uint8)
	fmt.Println(c, "-", d, "=", c-d)
}
