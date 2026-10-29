package gildedrose

type UpdatableItem interface {
	Update()
	Name() string
	SellIn() int
	Quality() int
}

func NewUpdatableItem(name string, sellIn, quality int) UpdatableItem {
	item := Item{
		name:    name,
		sellIn:  sellIn,
		quality: quality,
	}
	switch name {
	case "Aged Brie":
		return NewAgedBrie(item)
	case "Backstage passes to a TAFKAL80ETC concert":
		return NewBackstagePasses(item)
	case "Conjured":
		return NewConjured(item)
	case "Sulfuras, Hand of Ragnaros":
		return NewSulfuras(item)
	default:
		return NewNormal(item)
	}
}
