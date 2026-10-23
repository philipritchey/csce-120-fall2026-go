package main

import "fmt"

// I have an interface: Billable
// Billable types have a CostCents() method
type Billable interface {
	CostCents() int
}

// I have a type: ParkingPass
type ParkingPass struct {
	Hours int
}

// ParkingPass has a CostCents() method
// So ParkingPass is a Billable type (satisfies the Billable interface)
func (p ParkingPass) CostCents() int {
	return p.Hours * 200
}

// I have another type: MuseumTicket
type MuseumTicket struct {
	Age int
}

// MuseumTicket has a CostCents() method
// So MuseumTicket is a Billable type (satisfies the Billable interface)
func (m MuseumTicket) CostCents() int {
	if m.Age < 12 {
		return 500
	}
	return 1200
}

// Total computes the total cost in cents of a list of Billable items
func Total(items []Billable) int {
	total := 0
	for _, item := range items {
		total += item.CostCents()
	}
	return total
}

func main() {
	// items is a list of Billable values
	// ParkingPass and MuseumTicket are both Billable types
	// So we can put their values into the slice
	items := []Billable{
		ParkingPass{Hours: 3},
		MuseumTicket{Age: 10},
		MuseumTicket{Age: 30},
	}

	// And we can compute their total cost
	fmt.Println(Total(items))

	// This is nifty because Total() doesn't have to care what concrete types
	//   are in the slice, it only cares that they have the CostCents() method
	// Interfaces are behavioral contracts between functions and types
	//   function Foo: "I need these methods: X(), Y(), Z()"
	//   interface I: "Satisfying types must have these methods: X(), Y(), Z()"
	//   	function: "I accept values that satisfy interface I"
	//   types A,B: "I have these methods: X(), Y(), Z()"
	//   	interface: "You satisfy interface I"
	//   func Foo(value I) can be called with types A and B
	//     one definition
	//     Foo(A{})
	//     Foo(B{})
}
