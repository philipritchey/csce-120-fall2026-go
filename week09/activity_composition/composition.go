package main

import "fmt"

type PersonInfo struct {
	// TODO
	Name string
	ID   int
}

// TODO: BadgeLine
func (p PersonInfo) BadgeLine() string {
	return fmt.Sprintf("%s #%d", p.Name, p.ID)
}

type Firefighter struct {
	// TODO
	PersonInfo
	weapon string
}

type Paramedic struct {
	// TODO
	PersonInfo
	rating int
}

// TODO: String
func (f Firefighter) String() string {
	return fmt.Sprintf("%s wields a %s", f.BadgeLine(), f.weapon)
}

func (p Paramedic) String() string {
	return fmt.Sprintf("%s rated %d", p.BadgeLine(), p.rating)
}

// TODO(extra): BadgeHolder
type BadgeHolder interface {
	BadgeLine() string
}

func main() {
	// TODO
	f := Firefighter{
		Name:   "Alice",
		ID:     7,
		weapon: "axe",
	}
	p := Paramedic{
		Name:   "Bob",
		ID:     13,
		rating: 3,
	}
	fmt.Println(f)
	fmt.Println(p)

	a := []BadgeHolder{f, p}
	for _, b := range a {
		fmt.Println(b.BadgeLine())
	}
}
