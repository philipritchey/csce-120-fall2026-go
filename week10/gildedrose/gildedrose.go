package gildedrose

func UpdateQuality(items []UpdatableItem) {
	for _, item := range items {
		item.Update()
	}
}
