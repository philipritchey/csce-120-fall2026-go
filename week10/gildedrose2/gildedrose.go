package gildedrose

type Item struct {
	Name    string
	SellIn  int
	Quality int
}

func (item Item) GetName() string {
	return item.Name
}

func (item Item) GetSellIn() int {
	return item.SellIn
}

func (item Item) GetQuality() int {
	return item.Quality
}

func (item *Item) Update() {
	// do nothing, on purpose
}

type UpdatableItem interface {
	Update()
	GetName() string
	GetSellIn() int
	GetQuality() int
}

func NewUpdatableItem(name string, sellIn, quality int) UpdatableItem {
	item := Item{
		Name:    name,
		SellIn:  sellIn,
		Quality: quality,
	}
	switch name {
	case "Aged Brie":
		return NewAgedBrie(item)
	case "Sulfuras, Hand of Ragnaros":
		return NewSulfuras(item)
	case "Backstage passes to a TAFKAL80ETC concert":
		return NewBackstagePasses(item)
	case "Conjured":
		return NewConjured(item)
	default:
		return NewNormal(item)
	}
}

type Normal struct {
	Item
}

func NewNormal(item Item) *Normal {
	return &Normal{item}
}

func (item *Normal) Update() {
	item.SellIn--
	if item.Quality > 0 {
		item.Quality--
	}
	if item.SellIn < 0 && item.Quality > 0 {
		item.Quality--
	}
}

type AgedBrie struct {
	Item
}

func NewAgedBrie(item Item) *AgedBrie {
	return &AgedBrie{item}
}

func (item *AgedBrie) Update() {
	item.SellIn--
	if item.Quality < 50 {
		item.Quality++
	}
	if item.SellIn < 0 && item.Quality < 50 {
		item.Quality++
	}
}

type BackstagePasses struct {
	Item
}

func NewBackstagePasses(item Item) *BackstagePasses {
	return &BackstagePasses{item}
}

func (item *BackstagePasses) Update() {
	item.SellIn--
	if item.Quality < 50 {
		item.Quality++
	}
	if item.SellIn < 10 && item.Quality < 50 {
		item.Quality++
	}
	if item.SellIn < 5 && item.Quality < 50 {
		item.Quality++
	}
	if item.SellIn < 0 {
		item.Quality = 0
	}
}

type Sulfuras struct {
	Item
}

func NewSulfuras(item Item) *Sulfuras {
	return &Sulfuras{item}
}

type Conjured struct {
	Normal
}

func NewConjured(item Item) *Conjured {
	return &Conjured{Normal{item}}
}

func (item *Conjured) Update() {
	item.Normal.Update()
	item.SellIn++ // undo the day decrease once
	item.Normal.Update()
}

func UpdateQuality(items []UpdatableItem) {
	for _, item := range items {
		item.Update()
	}
}
