package main

import "math"

func Main() { // function main is undeclared in the main package
	var a int // declared and not used

	b := 7
	math.Sqrt(b) // cannot use X (value of type Y) as Z value

	c++ // undefined

	d := foo() // assignment mismatch: 1 value but functionName returns 2 values
}

func foo() (int, int) {
	return 1, 2
}
