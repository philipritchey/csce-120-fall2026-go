package main

type Snack struct {
	Name     string
	Calories int
}

// 1. localCalculation
// 2. saveSnackForLater
// 3. main and stack frames
// 4. escape analysis

// the "go:noinline" comments are there to make the memory analysis output easier to read

//go:noinline
func localCalculation() int {
	// cups and ounces are local variables.
	// They are only needed while localCalculation is running.
	// Since nothing outside the function needs their addresses,
	// they can usually stay on the stack or be optimized away.
	cups := 2
	ounces := cups * 8
	return ounces
}

//go:noinline
func saveSnackForLater() *Snack {
	// snack is created inside the function
	// it is a local variable
	snack := Snack{
		Name:     "Granola Bar",
		Calories: 180,
	}

	// but we return its address
	return &snack

	// Because something outside the function can still access snack,
	// Go cannot throw it away when the function ends.
	// So snack "escapes" to the heap
}

func main() {
	_ = localCalculation()

	savedSnack := saveSnackForLater()

	_ = savedSnack.Name
}

// Run with: go build -gcflags="-m" demo.go
// This asks the Go compiler to explain some of its memory decisions.

// Output should be (like):
// ./demo.go:23:2: moved to heap: snack

// This means the compiler decided snack cannot stay only in the function's stack frame.

// If time permits,
// Run with: go build -gcflags="-m -m" demo.go
// This asks the Go compiler give more a detailed explaination of its memory decisions.
