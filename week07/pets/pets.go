package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Pet struct {
	Name    string `json:"name"`
	Species string `json:"species"`
	Age     int    `json:"age"`
}

func main() {
	data, err := os.ReadFile("pets.json")
	if err != nil {
		log.Fatal(err)
	}
	var pets []Pet
	err = json.Unmarshal(data, &pets)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("All pets:")
	for _, pet := range pets {
		fmt.Printf("%s is a %d-year-old %s.\n", pet.Name, pet.Age, pet.Species)
	}

	fmt.Println()

	fmt.Println("Cats:")
	for _, pet := range pets {
		if pet.Species == "cat" {
			fmt.Println(pet.Name)
		}
	}

	fmt.Println()

	fmt.Println("Pets aged 5 or older:")
	for _, pet := range pets {
		if pet.Age >= 5 {
			fmt.Println(pet.Name)
		}
	}
}
