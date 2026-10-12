package main

import "fmt"

// Vector2 is a vector in 2D space
type Vector2 struct {
	x, y float64
}

// the string representation of the vector
func (v Vector2) String() string {
	return fmt.Sprintf("<%f, %f>", v.x, v.y)
}

// A method with a value receiver gets a copy
func (a Vector2) Add(b Vector2) Vector2 {
	//    ^ this is a copy of LHS of invocation v.Add(w)
	return Vector2{a.x + b.x, a.y + b.y}
}

// A method with a pointer receiver can change the original struct value
func (a *Vector2) Update(b Vector2) {
	//    ^ this is address of LHS of invocation v.Update(w)
	a.x += b.x
	a.y += b.y
}

// Pointer receiver syntax uses *Type

func main() {
	// Go still lets us call pointer receiver methods with dot notation
	{
		v1 := Vector2{3, 4}
		v2 := Vector2{-4, 3}

		v1.Update(v2)    // OK -> Update has pointer receiver, &v1 will be used as receiver
		(&v1).Update(v2) // equivalent
	}

	// We can also call value receiver methods on pointers...
	{
		v1 := &Vector2{3, 4} // pointer
		v2 := Vector2{-4, 3} // value

		v1.Update(v2)    // OK -> Update has pointer receiver, v1 is pointer
		fmt.Println(v1)  // also OK -> String has value receiver, *v1 is used as receiver
		fmt.Println(*v1) // equivalent
	}

	// Go is just being kind
	// Go "spots" you the & or * to help you keep the code looking clean (and easier to type)
	//   "syntactic sugar"
	// You cannot have 2 methods with same name, one with a value receiver and the other with a pointer receiver
	//   That is ambiguous, Go won't allow it
}
