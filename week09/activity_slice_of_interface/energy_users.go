package main

import "fmt"

type EnergyUser interface {
	KWhPerDay() float64
}

type WattHoursPerDay struct {
	Watts int
	Hours int
}

type Lamp struct {
	WattHoursPerDay
}

type Laptop struct {
	WattHoursPerDay
}

func (wh WattHoursPerDay) KWhPerDay() float64 {
	return float64(wh.Watts*wh.Hours) / 1000.0
}

func TotalKWh(users []EnergyUser) float64 {
	total := 0.0
	for _, u := range users {
		total += u.KWhPerDay()
	}
	return total
}

func main() {
	users := []EnergyUser{
		Lamp{
			Watts: 60,
			Hours: 12,
		}, // 0.72 KWh
		Laptop{
			Watts: 400,
			Hours: 3,
		}, // 1.2 KWh
		Lamp{
			Watts: 9,
			Hours: 17,
		}, // 0.153 KWh
	}
	//   0.720
	//   1.200
	// + 0.153
	// -------
	//   2.073
	fmt.Println("total =", TotalKWh(users)) // should be 2.073
}
