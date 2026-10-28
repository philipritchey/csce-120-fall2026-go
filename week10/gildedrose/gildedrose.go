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

type UpdatableItem interface {
	Update()
	Name() string
	SellIn() int
	Quality() int
}

type Normal struct {
	Item
}

type AgedBrie struct {
	Item
}

type BackstagePasses struct {
	Item
}

type Sulfuras struct {
	Item
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

func (item *AgedBrie) Update() {
	item.sellIn--
	if item.quality < 50 {
		item.quality++
	}
	if item.sellIn < 0 && item.quality < 50 {
		item.quality++
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

func (item *Sulfuras) Update() {
	// do nothing!
}

func UpdateQuality(items []*Item) {
	for _, item := range items {
		var t UpdatableItem
		switch item.Name() {
		case "Aged Brie":
			t = &AgedBrie{Item: *item}
		case "Backstage passes to a TAFKAL80ETC concert":
			t = &BackstagePasses{Item: *item}
		case "Sulfuras, Hand of Ragnaros":
			t = &Sulfuras{Item: *item}
		default:
			t = &Normal{Item: *item}
		}
		t.Update()
		item.sellIn, item.quality = t.SellIn(), t.Quality()
	}
}
