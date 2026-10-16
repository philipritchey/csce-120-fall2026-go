package main

type Student struct {
	Name string
}

//go:noinline
func main() {
	s := &Student{Name: "Alice"}

	students := []*Student{}
	students = append(students, s)

	_ = students
}
