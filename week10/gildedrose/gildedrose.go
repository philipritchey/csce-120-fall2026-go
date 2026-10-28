package gildedrose

type Item struct {
	Name    string
	SellIn  int
	Quality int
}

func UpdateNormal(item *Item) {
	item.SellIn--
	if item.Quality > 0 {
		item.Quality--
	}
	if item.SellIn < 0 && item.Quality > 0 {
		item.Quality--
	}
}

func UpdateAgedBrie(item *Item) {
	item.SellIn--
	if item.Quality < 50 {
		item.Quality++
	}
	if item.SellIn < 0 && item.Quality < 50 {
		item.Quality++
	}
}

func UpdateBackstagePasses(item *Item) {
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

func UpdateSulfuras(item *Item) {
	// do nothing!
}

func UpdateQuality(items []*Item) {
	for _, item := range items {
		switch item.Name {
		case "Aged Brie":
			UpdateAgedBrie(item)
		case "Backstage passes to a TAFKAL80ETC concert":
			UpdateBackstagePasses(item)
		case "Sulfuras, Hand of Ragnaros":
			UpdateSulfuras(item)
		default:
			UpdateNormal(item)
		}
	}
}
