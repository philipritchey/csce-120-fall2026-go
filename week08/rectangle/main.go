package main

import "fmt"

func main() {
	r1 := Rectangle{Width: 4, Height: 6}
	r2 := Rectangle{Width: 2.5, Height: 3}

	fmt.Println("Rectangle:", r1)
	fmt.Println("Area:", r1.Area())
	fmt.Println("Perimeter:", r1.Perimeter())

	fmt.Println()

	fmt.Println("Rectangle:", r2)
	fmt.Println("Area:", r2.Area())
	fmt.Println("Perimeter:", r2.Perimeter())

	fmt.Println()

	c := Circle{1}
	fmt.Println("Circle:", c)
	fmt.Println("Area:", c.Area())
	fmt.Println("Circumference", c.Circumference())
}
