package gildedrose

type Sulfuras struct {
	Item
}

func NewSulfuras(name string, sellIn, quality int) *Sulfuras {
	return &Sulfuras{
		name:    name,
		sellIn:  sellIn,
		quality: quality,
	}
}

func (item *Sulfuras) Update() {
	// do nothing!
}
