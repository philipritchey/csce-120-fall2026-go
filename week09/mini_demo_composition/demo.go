package main

import "fmt"

type ContactInfo struct {
	Email string
}

func (c ContactInfo) ContactLine() string {
	return "Email: " + c.Email
}

// Volunteer embeds ContactInfo
// "composition" := building a bigger type out of smaller types
// where inheritance says "X is a Y",
// composition says "X has a Y"
// Volunteer has ContactInfo
// embedding is Go's shortcut syntax for composition
type Volunteer struct {
	ContactInfo
	Hours int
}

// Donor has ContactInfo, too
type Donor struct {
	ContactInfo
	Amount int
}

// Composition is useful for sharing data (fields) and behavior
// Both Volunteer and Donor have an Email field (because they have ContactInfo)
// Both Volunteer and Donor have a ContactLine() method (because they have ContactInfo)

func main() {
	// explicit declaration of field type and value (can be removed)
	v := Volunteer{
		ContactInfo: ContactInfo{Email: "v@example.com"},
		Hours:       12,
	}

	// direct assignment of Email through composition/embedding
	// Donor has an Email field becasue Donor has a ContactInfo, which has Email
	d := Donor{
		Email:  "d@example.com",
		Amount: 50,
	}

	// Volunteers and Donors have a ContactLine() method through ContactInfo
	// NOT inheritance: Volunteer is not derived from ContactInfo, not "is a"
	// not an interface: Volunteer does not implement/satisfy ContactInfo (which isn't an interface anyways)
	// composition: Volunteer has a ContactInfo, Go's embedding takes a shortcut
	// embedding helps you reuse fields and methods without copying code
	fmt.Println(v.ContactLine())
	fmt.Println(d.ContactLine())
	// equivalent, but longer, syntax
	fmt.Println(v.ContactInfo.ContactLine())

	// can access fields directly
	fmt.Println(v.Email) // short version using embedding shortcut
	fmt.Println(v.ContactInfo.Email)

	// composition and interfaces work together as two parts of the same design idea:
	//   build programs out of small pieces that provide useful behavior
	// composition helps a type gain behavior
	// interfaces let code use any type that has that behavior
	// composition is how you assemble types
	// interfaces are how you describe and use their behavior
}
