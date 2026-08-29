package main

import "fmt"

type PhoneNumber int
type ZipCode int

func MakeCall(n PhoneNumber) {
	fmt.Println("calling", n)
}
func SendMail(z ZipCode) {
	fmt.Println("mailing", z)
}

func main() {
	var phoneNumber PhoneNumber = 8675309
	var zipCode ZipCode = 77845
	MakeCall(phoneNumber)
	SendMail(zipCode)
	MakeCall(90210)      // compiler allows, but "clearly" wrong
	SendMail(9798453510) // compiler allows, but "clearly" wrong
	// MakeCall(z) // compiler won't allow (good)
	// SendMail(p) // compiler won't allow (good)
}
