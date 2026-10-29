package gildedrose

type Conjured struct {
	Normal
}

func NewConjured(name string, sellIn, quality int) *Conjured {
	return &Conjured{
		name:    name,
		sellIn:  sellIn,
		quality: quality,
	}
}

func (item *Conjured) Update() {
	// Conjured degrades twice as fast as Normal...
	item.Normal.Update()
	item.Normal.sellIn++ // "undo" the day decrement once
	item.Normal.Update()
}
