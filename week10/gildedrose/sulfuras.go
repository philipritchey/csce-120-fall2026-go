package gildedrose

type Sulfuras struct {
	Item
}

func NewSulfuras(item Item) *Sulfuras {
	return &Sulfuras{item}
}
