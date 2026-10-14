package main

import (
	"bufio"
	"fmt"
	"os"
)

type Artifact struct {
	Name string
	Year int
}

func main() {
	// Allocation using `new`
	// new(Artifact) creates a zero-valued Artifact and returns a pointer to it
	artifact1 := new(Artifact)
	fmt.Println("artifact1:", artifact1)

	pause()

	// after a value is allocated, we can access (read/write) it
	artifact1.Name = "Clay Bowl"
	artifact1.Year = 1200
	fmt.Println("artifact1:", artifact1)

	pause()

	// Allocation using composite literal
	// &Artifact{...} creates an initialized Artifact and returns its address.
	artifact2 := &Artifact{
		Name: "Bronze Coin",
		Year: 450,
	}
	fmt.Println("artifact2:", artifact2)

	pause()

	// Allocation using `make`
	// make is used for built-in data structures like slices and maps.
	museumCase := make([]*Artifact, 0)
	rooms := make(map[string]string)

	// append may re-allocate the slice (to be bigger)
	museumCase = append(museumCase, artifact1)
	museumCase = append(museumCase, artifact2)

	rooms["Clay Bowl"] = "Ancient History"
	rooms["Bronze Coin"] = "Economics and Trade"

	for _, artifact := range museumCase {
		fmt.Printf("%s, year %d, room: %s\n",
			artifact.Name,
			artifact.Year,
			rooms[artifact.Name],
		)
	}

	pause()

	// recall: we can also allocate slices and maps using literals
	sliceOfInt := []int{1, 2, 3}
	mapStringToInt := map[string]int{"x": 1, "y": 2, "z": 3}
	fmt.Println("sliceOfInt:", sliceOfInt)
	fmt.Println("mapStringToInt:", mapStringToInt)

	pause()

	// `var s []int` allocates nothing
	// s is nil
	// but append can still be used
	var s []int
	fmt.Println("s:", s, "is nil:", s == nil)
	s = append(s, 1)
	fmt.Println("s:", s, "is nil:", s == nil)
}

func pause() {
	fmt.Printf("\n...paused...\n")
	s := bufio.NewScanner(os.Stdin)
	s.Scan()
}
