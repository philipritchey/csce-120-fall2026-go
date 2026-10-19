package main

import "fmt"

type GardenPlant struct {
	Name           string
	WaterEveryDays int
	Sunlight       string
}

func NewGardenPlant(name string, water int, sunlight string) GardenPlant {
	return GardenPlant{
		Name:           name,
		WaterEveryDays: water,
		Sunlight:       sunlight,
	}
}

func (p GardenPlant) CareInstructions() string {
	return fmt.Sprintf("Water %s every %d days, needs %s", p.Name, p.WaterEveryDays, p.Sunlight)
}

func main() {
	strawberry := NewGardenPlant("strawberry", 4, "partial sun")
	dragonfruit := NewGardenPlant("dragonfruit", 14, "lots of sun")
	fmt.Println(strawberry.CareInstructions())
	fmt.Println(dragonfruit.CareInstructions())
}
