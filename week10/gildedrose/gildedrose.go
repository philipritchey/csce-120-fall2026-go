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
	for i := 0; i < len(items); i++ {
		if items[i].Name != "Aged Brie" &&
			items[i].Name != "Backstage passes to a TAFKAL80ETC concert" &&
			items[i].Name != "Sulfuras, Hand of Ragnaros" {
			UpdateNormal(items[i])
			continue
		}
		if items[i].Name == "Aged Brie" {
			UpdateAgedBrie(items[i])
			continue
		}
		if items[i].Name == "Backstage passes to a TAFKAL80ETC concert" {
			UpdateBackstagePasses(items[i])
			continue
		}
		if items[i].Name == "Sulfuras, Hand of Ragnaros" {
			UpdateSulfuras(items[i])
			continue
		}
	}
}
