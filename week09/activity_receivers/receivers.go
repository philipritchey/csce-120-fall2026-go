package main

import "fmt"

type StockItem struct {
	Name     string
	Quantity int
}

func (s *StockItem) Add(count int) {
	if count <= 0 {
		return
	}
	s.Quantity += count
}

func (s *StockItem) Remove(count int) {
	if count <= 0 || count > s.Quantity {
		return
	}
	s.Quantity -= count
}

func (s *StockItem) Rename(newName string) {
	s.Name = newName
}

type Adjustable interface {
	Add(int)
	Remove(int)
}

func Restock(a Adjustable) {
	a.Add(10)
}

// Anything that has Add and Remove satisfies Adjustable
type Set struct {
	Values map[int]struct{}
}

func (s *Set) Add(value int) {
	s.Values[value] = struct{}{}
}

func (s *Set) Remove(value int) {
	delete(s.Values, value)
}

func main() {
	item := StockItem{Name: "markers", Quantity: 10}

	item.Add(5)
	item.Remove(3)

	fmt.Println(item.Quantity) // should print 12

	item.Rename("pens")
	fmt.Println(item.Name) // should print pens

	item.Add(-1)
	fmt.Println(item.Quantity) // should print 12
	item.Remove(-1)
	fmt.Println(item.Quantity) // should print 12
	item.Remove(500)
	fmt.Println(item.Quantity) // should print 12
	item.Remove(12)
	fmt.Println(item.Quantity) // should print 0

	Restock(&item)             // note: must be pointer, not value (because mutation)
	fmt.Println(item.Quantity) // should print 10

	// a bit awkward, maybe...
	// but Set has Add and Remove
	// so Set satsifies Adjustable
	// so Set can be passed to Restock
	set := &Set{
		Values: make(map[int]struct{}),
	}
	set.Add(1)
	set.Add(2)
	set.Add(3)
	fmt.Println(set) // should have {1 2 3}
	set.Remove(3)
	fmt.Println(set) // should have {1 2}
	Restock(set)
	fmt.Println(set) // should have {1 2 10}
}
