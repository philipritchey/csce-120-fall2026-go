package gildedrose

type AgedBrie struct {
	Item
}

func NewAgedBrie(name string, sellIn, quality int) *AgedBrie {
	return &AgedBrie{
		name:    name,
		sellIn:  sellIn,
		quality: quality,
	}
}

func (item *AgedBrie) Update() {
	item.sellIn--
	if item.quality < 50 {
		item.quality++
	}
	if item.sellIn < 0 && item.quality < 50 {
		item.quality++
	}
}
