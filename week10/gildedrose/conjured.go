package gildedrose

type Conjured struct {
	Normal
}

func NewConjured(item Item) *Conjured {
	return &Conjured{Normal{item}}
}

func (item *Conjured) Update() {
	// Conjured degrades twice as fast as Normal...
	item.Normal.Update()
	item.Normal.sellIn++ // "undo" the day decrement once
	item.Normal.Update()
}
