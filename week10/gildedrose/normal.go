package gildedrose

type Normal struct {
	Item
}

func NewNormal(item Item) *Normal {
	return &Normal{item}
}

func (item *Normal) Update() {
	item.sellIn--
	if item.quality > 0 {
		item.quality--
	}
	if item.sellIn < 0 && item.quality > 0 {
		item.quality--
	}
}
