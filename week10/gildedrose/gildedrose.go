package gildedrose

type Item struct {
	Name    string
	SellIn  int
	Quality int
}

type Updatable interface {
	Update()
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
	item.SellIn--
	if item.Quality > 0 {
		item.Quality--
	}
	if item.SellIn < 0 && item.Quality > 0 {
		item.Quality--
	}
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

func (item *Sulfuras) Update() {
	// do nothing!
}

func UpdateQuality(items []*Item) {
	for _, item := range items {
		switch item.Name {
		case "Aged Brie":
			t := AgedBrie{Item: *item}
			t.Update()
			item.SellIn, item.Quality = t.SellIn, t.Quality
		case "Backstage passes to a TAFKAL80ETC concert":
			t := BackstagePasses{Item: *item}
			t.Update()
			item.SellIn, item.Quality = t.SellIn, t.Quality
		case "Sulfuras, Hand of Ragnaros":
			t := Sulfuras{Item: *item}
			t.Update()
			item.SellIn, item.Quality = t.SellIn, t.Quality
		default:
			t := Normal{Item: *item}
			t.Update()
			item.SellIn, item.Quality = t.SellIn, t.Quality
		}
	}
}
