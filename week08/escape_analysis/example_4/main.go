package main

type Point struct {
	X int
	Y int
}

func main() {
	p1 := new(Point)
	p2 := &Point{}

	_ = p1
	_ = p2
}
