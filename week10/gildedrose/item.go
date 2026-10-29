package gildedrose

type Item struct {
	name    string
	sellIn  int
	quality int
}

func (i *Item) Update() {
	// do nothing
}

func (i Item) Name() string {
	return i.name
}

func (i Item) SellIn() int {
	return i.sellIn
}

func (i Item) Quality() int {
	return i.quality
}
