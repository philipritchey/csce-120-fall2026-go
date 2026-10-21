package main

import "fmt"

type Billable interface {
	CostCents() int
}

type ParkingPass struct {
	Hours int
}

func (p ParkingPass) CostCents() int {
	return p.Hours * 200
}

type MuseumTicket struct {
	Age int
}

func (m MuseumTicket) CostCents() int {
	if m.Age < 12 {
		return 500
	}
	return 1200
}

func Total(items []Billable) int {
	total := 0
	for _, item := range items {
		total += item.CostCents()
	}
	return total
}

func main() {
	items := []Billable{
		ParkingPass{Hours: 3},
		MuseumTicket{Age: 10},
		MuseumTicket{Age: 30},
	}

	fmt.Println(Total(items))
}
