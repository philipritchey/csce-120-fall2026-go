package main

import "fmt"

type Voter struct {
	Name         string
	County       string
	Age          int
	IsRegistered bool
}

func main() {
	voters := []Voter{
		{"Alice", "Brazos", 42, true},
		{"Bob", "Harris", 67, false},
		{"Carol", "Waller", 17, false},
		{"Dave", "Travis", 19, true},
	}
	fmt.Println("All voters:")
	for _, voter := range voters {
		var isOrIsnt string
		if voter.IsRegistered {
			isOrIsnt = "is"
		} else {
			isOrIsnt = "is not"
		}
		fmt.Printf("%s is %d years old and %s registered to vote in %s county\n", voter.Name, voter.Age, isOrIsnt, voter.County)
	}

	fmt.Println()

	fmt.Println("Unregistered voters with age >= 18:")
	for _, voter := range voters {
		if !voter.IsRegistered && voter.Age >= 18 {
			fmt.Printf("%s in %s county\n", voter.Name, voter.County)
		}
	}
}
