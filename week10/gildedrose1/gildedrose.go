package gildedrose

type Item struct {
	name    string
	sellIn  int
	quality int
}

func (item *Item) Update() {
	// intentionally empty
}

func (item Item) Name() string {
	return item.name
}

func (item Item) SellIn() int {
	return item.sellIn
}

func (item Item) Quality() int {
	return item.quality
}

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
	case "Sulfuras, Hand of Ragnaros":
		return NewSulfuras(name, sellIn, quality)
	case "Backstage passes to a TAFKAL80ETC concert":
		return NewBackstagePasses(name, sellIn, quality)
	case "Conjured":
		return NewConjured(name, sellIn, quality)
	default:
		return NewNormal(name, sellIn, quality)
	}
}

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

type Sulfuras struct {
	Item
}

func NewSulfuras(name string, sellIn, quality int) *Sulfuras {
	return &Sulfuras{
		name:    name,
		sellIn:  sellIn,
		quality: quality,
	}
}

func (item *Sulfuras) Update() {
	// do nothing!
}

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
	item.Normal.Update()
	item.sellIn++ // undo the day decrement once
	item.Normal.Update()
}

func UpdateQuality(items []UpdatableItem) {
	for _, item := range items {
		item.Update()
	}
}
