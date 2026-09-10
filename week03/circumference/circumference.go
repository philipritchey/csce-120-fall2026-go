package main

import (
	"fmt"
	"math"
)

func main() {
	radius, err := readRadius()
	if err != nil {
		fmt.Printf("invalid: %v\n", err)
		return
	}
	printCircleInfo(radius)
}

// readRadius reads a non-negative floating point number from standard input.
// It returns the value read or an error explaining what went wrong.
func readRadius() (float64, error) {
	var radius float64
	_, err := fmt.Scan(&radius)
	if err != nil {
		return 0, err
	}
	if radius < 0 {
		return 0, fmt.Errorf("radius < 0")
	}
	return radius, nil
}

// printCiricleInfo prints information about a circle to standard output.
// The information printed is the radius, diameter, circumference, and area.
func printCircleInfo(radius float64) {
	fmt.Println("Circle Info")
	fmt.Println("       Radius:", radius)
	fmt.Println("     Diameter:", diameter(radius))
	fmt.Println("Circumference:", circumference(radius))
	fmt.Println("         Area:", area(radius))
}

// diameter returns the diameter of a circle.
//
// Parameters:
//
//	radius: the radius of the circle
//
// Returns:
//
//	The diameter of the circle
//
// Preconditions:
//
//	radius should be non-negative
//
// Postconditions:
//
//	The returned value is 2 * radius
func diameter(radius float64) float64 {
	return 2 * radius
}

// circumference returns the circumference of a circle.
//
// Parameters:
//
//	radius: the radius of the circle
//
// Returns:
//
//	The circumference of the circle.
//
// Preconditions:
//
//	radius should be positive.
//
// Postconditions:
//
//	The returned value equals 2 * math.Pi * radius.
func circumference(radius float64) float64 {
	return 2 * math.Pi * radius
}

// area returns the area of a circle.
//
// Parameters:
//
//	radius: the radius of the circle
//
// Returns:
//
//	The area of the circle.
//
// Preconditions:
//
//	None
//
// Postconditions:
//
//	The returned value equals math.Pi * radius * radius.
func area(radius float64) float64 {
	return math.Pi * radius * radius
}
