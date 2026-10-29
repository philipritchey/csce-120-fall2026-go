package gildedrose

type AgedBrie struct {
	Item
}

func NewAgedBrie(item Item) *AgedBrie {
	return &AgedBrie{item}
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
