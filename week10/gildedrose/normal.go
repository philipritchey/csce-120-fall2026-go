package gildedrose

type Normal struct {
	Item
}

func NewNormal(name string, sellIn, quality int) *Normal {
	return &Normal{
		name:    name,
		sellIn:  sellIn,
		quality: quality,
	}
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
