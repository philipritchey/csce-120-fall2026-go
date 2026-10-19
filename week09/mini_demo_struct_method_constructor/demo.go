package main

import "fmt"

// a struct groups related data
type Package struct {
	Recipient string
	WeightOz  int
}

// constructor := a function that creates values of a certain type
// NewTypeName is the conventional name for a constructor in Go
// return a value if the type is small and immutable
// return a pointer if the type is large or mutable (most common)
func NewPackage(recipient string, weightOz int) Package {
	return Package{
		Recipient: recipient,
		WeightOz:  weightOz,
	}
}

// values of type Package know how to produce a label
// use pointer receiver if method needs to modify the receiver
// otherwise, use value receiver
// be consistent: all pointer or all value
// when in doubt: pointer
func (p Package) Label() string {
	return fmt.Sprintf("Ship to %s: %d oz", p.Recipient, p.WeightOz)
}

func main() {
	p := NewPackage("Avery", 12)
	fmt.Println(p.Label())

	// q := &NewPackage("Bill", 42)  // invalid because return value is temporary -> no address
	q := NewPackage("Bill", 42)
	qq := &q
	fmt.Println(qq.Label())

	// this works, but it would be better just to have constructor return a pointer
	r := new(NewPackage("Cat", 23))
	fmt.Println(r.Label())

}
