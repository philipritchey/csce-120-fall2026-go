package gildedrose

type BackstagePasses struct {
	Item
}

func NewBackstagePasses(name string, sellIn, quality int) *BackstagePasses {
	return &BackstagePasses{
		name:    name,
		sellIn:  sellIn,
		quality: quality,
	}
}

func (item *BackstagePasses) Update() {
	item.sellIn--
	if item.quality < 50 {
		item.quality++
	}
	if item.sellIn < 10 && item.quality < 50 {
		item.quality++
	}
	if item.sellIn < 5 && item.quality < 50 {
		item.quality++
	}
	if item.sellIn < 0 {
		item.quality = 0
	}
}
