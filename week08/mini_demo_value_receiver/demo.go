package main

import (
	"fmt"
	"math"
)

// A method is a function attached to a type
// Vector2 is a vector in 2D space
type Vector2 struct {
	x, y float64
}

// function syntax
// add 2 vectors in 2D space
func Add(a, b Vector2) Vector2 {
	return Vector2{a.x + b.x, a.y + b.y}
}

// method syntax
func (a Vector2) Add(b Vector2) Vector2 {
	return Vector2{a.x + b.x, a.y + b.y}
}

// The receiver tells Go to which type the method belongs
// method syntax
// func (a Vector2) Add(b Vector2) Vector2 {...}
//      ^^^^^^^^^^^ this part is the "receiver"

func main() {
	// Methods are called with dot notation, like value.Method()
	v1 := Vector2{3, 4}
	v2 := Vector2{-4, 3}

	v3 := v1.Add(v2)
	fmt.Println(v1, "+", v2, "=", v3)
}

// Value receiver methods are useful when the method only needs to read data, not change it
// the length of the vector
func (v Vector2) Magnitude() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

// the string representation of the vector
func (v Vector2) String() string {
	return fmt.Sprintf("<%f, %f>", v.x, v.y)
}
