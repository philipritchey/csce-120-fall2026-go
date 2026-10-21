package main

import "fmt"

type ScoreCard struct {
	Points int
}

// value receiver: receives a copy
// update only affects the copy, not the original
func (s ScoreCard) AddWrong(points int) {
	s.Points += points
}

// pointer receiver: receives a pointer
// update affects the original
// that v (*) is the only difference (easy to miss when debugging)
func (s *ScoreCard) AddRight(points int) {
	s.Points += points
}

// we'll come back to these at the end...
type PointAdder interface {
	AddWrong(int)
	AddRight(int)
}

func AddPoints(a PointAdder, points int) {
	a.AddWrong(points)
	a.AddRight(points)
}

func main() {
	card := ScoreCard{Points: 10}

	card.AddWrong(5)
	fmt.Println(card.Points) // expected 15, got 10 (because value receiver)

	card.AddRight(5)
	fmt.Println(card.Points) // 15 (because pointer receiver)

	// choosing between a value receiver and a pointer receiver is a design choice
	//   about what the method means
	// if the method only observes the value, can use a value receiver
	// if the method should change the object, must use a pointer receiver
	// connection to interfaces
	//   interfaces describe what a type can do
	//   e.g.
	//     SetName() sounds like it changes the value, so should have a pointer receiver
	//     GetName() sounds like it only reads the value, so should have a value receiver

	// receiver choice is part of the contract your type offers
	//   a value receiver says "I can do this without changing myself"
	//   a pointer receiver says "This behavior may change me, so I need access to the original object"
	// good design means the receiver matches the meaning of the method

	// it is legal Go to mix value and pointer receivers on the same type,
	//   but idiomatic Go favors consistency
	// the main reason consistency matters is method sets
	// e.g.
	//   if we define
	//     func (i Item) GetTitle() string
	//     func (i *Item) SetTitle(title string)
	//   then
	//     *Item has BOTH the value-receiver and pointer-receiver methods
	//     Item has ONLY the value-receiver method
	//   so if an interface requires SetTitle, then usually *Item, not Item, satisfies that interface

	// card is a value so it does NOT satisfy the PointAdder interface
	//   which requires AddRight
	//   and AddRight has a pointer receiver
	// AddPoints(card, 5)
	// but &card does
	AddPoints(&card, 5)
	fmt.Println(card.Points)
}
