package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Why use a struct?
	// Real-world things usually need more than one value to describe them
	// Example: A student might have a name, a major, a year, and a GPA
	// We could store those separately, but it is sometimes easier to group them together
	name := "Alice"
	major := "Computer Science"
	year := 1
	gpa := 3.8
	fmt.Println(name, major, year, gpa)
	// These variables all describe the same student,
	// but the language does not know they belong together

	pause()

	// A struct would let us create one value that contains all of this related information
	type Student struct {
		Name  string
		Major string
		Year  int
		GPA   float64
	}
	// `type Student struct` means we are defining:
	// a new `type`
	// named `Student`
	// which is a `struct` with four fields
	// each field has a name and a type
	// every Student value contains all four fields

	// This creates one Student value:
	student1 := Student{
		Name:  "Alice",
		GPA:   3.8,
		Year:  1,
		Major: "Computer Science",
	}
	fmt.Println(student1)
	// the field names on the left match the fields from the struct definition
	// when names are given, the order of the values does not matter
	// `Name: "Alice"` means put `"Alice"` in the `Name` field
	// without names, the order is expected to match the order in the type definition
	// e.g. Student{"Alice", "Computer Science", 1, 3.8}
	// fields not given values will be initialized to their type's zero value

	pause()

	// Go converts a struct to a string by converting each field value to a string
	fmt.Println("student1:", student1)

	pause()

	// We access fields with dot notation
	// `student1.Name` means get the `Name` field from the `student1` struct value
	// The dot operator lets us access one piece of data (field)
	// inside a struct value (of potentially many fields)
	fmt.Println("Name:", student1.Name)
	fmt.Println("Major:", student1.Major)
	fmt.Println("Year:", student1.Year)
	fmt.Println("GPA:", student1.GPA)

	pause()

	fmt.Printf("%s is a %s major in year %d with a GPA of %.2f.\n",
		student1.Name,
		student1.Major,
		student1.Year,
		student1.GPA,
	)

	pause()

	// This creates another Student value:
	student2 := Student{
		Name:  "Bob",
		Major: "Biology",
		Year:  2,
		GPA:   3.2,
	}

	// Now we have 2 Student values
	fmt.Printf("%s is a %s major.\n", student1.Name, student1.Major)
	fmt.Printf("%s is a %s major.\n", student2.Name, student2.Major)

	pause()

	// Student is a type we can reuse
	// We can make a slice of Student to store our ever-growing list of students
	students := []Student{student1, student2}
	fmt.Println(students)

	// Recap
	// A struct lets us group related pieces of data into one custom reusable type
	// Each named piece of data inside a struct is called a field
	// We can create struct values using struct literals
	// We can access fields using dot notation
}

func pause() {
	fmt.Printf("\n...paused...\n")
	s := bufio.NewScanner(os.Stdin)
	s.Scan()
}
