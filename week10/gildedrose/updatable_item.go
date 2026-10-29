package gildedrose

type UpdatableItem interface {
	Update()
	Name() string
	SellIn() int
	Quality() int
}

func NewUpdatableItem(name string, sellIn, quality int) UpdatableItem {
	switch name {
	case "Aged Brie":
		return NewAgedBrie(name, sellIn, quality)
	case "Backstage passes to a TAFKAL80ETC concert":
		return NewBackstagePasses(name, sellIn, quality)
	case "Conjured":
		return NewConjured(name, sellIn, quality)
	case "Sulfuras, Hand of Ragnaros":
		return NewSulfuras(name, sellIn, quality)
	default:
		return NewNormal(name, sellIn, quality)
	}
}
