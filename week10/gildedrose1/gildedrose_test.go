package gildedrose

import (
	"fmt"
	"testing"
)

func TestUpdateQuality(t *testing.T) {
	type status struct {
		sellIn  int
		quality int
	}
	testCases := []struct {
		desc     string
		name     string
		initial  status
		expected status
	}{
		{
			desc:     "normal item",
			name:     "normal",
			initial:  status{sellIn: 10, quality: 20},
			expected: status{sellIn: 9, quality: 19},
		},
		{
			desc:     "normal item just before expiration",
			name:     "normal",
			initial:  status{sellIn: 1, quality: 20},
			expected: status{sellIn: 0, quality: 19},
		},
		{
			desc:     "normal item at expiration",
			name:     "normal",
			initial:  status{sellIn: 0, quality: 19},
			expected: status{sellIn: -1, quality: 17},
		},
		{
			desc:     "normal item past expiration has quality not less than 0",
			name:     "normal",
			initial:  status{sellIn: -1, quality: 1},
			expected: status{sellIn: -2, quality: 0},
		},
		{
			desc:     "aged brie before expiration",
			name:     "Aged Brie",
			initial:  status{sellIn: 10, quality: 20},
			expected: status{sellIn: 9, quality: 21},
		},
		{
			desc:     "aged brie before expiration has quality not more than 50",
			name:     "Aged Brie",
			initial:  status{sellIn: 10, quality: 50},
			expected: status{sellIn: 9, quality: 50},
		},
		{
			desc:     "aged brie after expiration improves twice as fast",
			name:     "Aged Brie",
			initial:  status{sellIn: 0, quality: 30},
			expected: status{sellIn: -1, quality: 32},
		},
		{
			desc:     "aged brie just before expiration improves like normal",
			name:     "Aged Brie",
			initial:  status{sellIn: 1, quality: 30},
			expected: status{sellIn: 0, quality: 31},
		},
		{
			desc:     "aged brie after expiration has quality not more than 50",
			name:     "Aged Brie",
			initial:  status{sellIn: -2, quality: 50},
			expected: status{sellIn: -3, quality: 50},
		},
		{
			desc:     "Sulfuras never has to be sold and never degrades in quality",
			name:     "Sulfuras, Hand of Ragnaros",
			initial:  status{sellIn: 10, quality: 80},
			expected: status{sellIn: 10, quality: 80},
		},
		{
			desc:     "Sulfuras after expiration never has to be sold and never degrades in quality",
			name:     "Sulfuras, Hand of Ragnaros",
			initial:  status{sellIn: -1, quality: 80},
			expected: status{sellIn: -1, quality: 80},
		},
		{
			desc:     "Backstage passes with more than 10 days improve by 1",
			name:     "Backstage passes to a TAFKAL80ETC concert",
			initial:  status{sellIn: 11, quality: 30},
			expected: status{sellIn: 10, quality: 31},
		},
		{
			desc:     "Backstage passes with 6 to 10 days improve by 2",
			name:     "Backstage passes to a TAFKAL80ETC concert",
			initial:  status{sellIn: 6, quality: 30},
			expected: status{sellIn: 5, quality: 32},
		},
		{
			desc:     "Backstage passes with 1 - 5 days improve by 3",
			name:     "Backstage passes to a TAFKAL80ETC concert",
			initial:  status{sellIn: 1, quality: 30},
			expected: status{sellIn: 0, quality: 33},
		},
		{
			desc:     "Backstage passes expired have quality 0",
			name:     "Backstage passes to a TAFKAL80ETC concert",
			initial:  status{sellIn: 0, quality: 30},
			expected: status{sellIn: -1, quality: 0},
		},
		{
			desc:     "Backstage passes quality never less than 0",
			name:     "Backstage passes to a TAFKAL80ETC concert",
			initial:  status{sellIn: -1, quality: 0},
			expected: status{sellIn: -2, quality: 0},
		},
		{
			desc:     "Backstage passes quality never more than 50 when far away",
			name:     "Backstage passes to a TAFKAL80ETC concert",
			initial:  status{sellIn: 20, quality: 50},
			expected: status{sellIn: 19, quality: 50},
		},
		{
			desc:     "Backstage passes quality never more than 50 when near",
			name:     "Backstage passes to a TAFKAL80ETC concert",
			initial:  status{sellIn: 8, quality: 50},
			expected: status{sellIn: 7, quality: 50},
		},
		{
			desc:     "Backstage passes quality never more than 50 when very soon",
			name:     "Backstage passes to a TAFKAL80ETC concert",
			initial:  status{sellIn: 3, quality: 50},
			expected: status{sellIn: 2, quality: 50},
		},
		{
			desc:     "Conjured unexpired item degrdaes twice as fast as normal unexpired",
			name:     "Conjured",
			initial:  status{sellIn: 10, quality: 8},
			expected: status{sellIn: 9, quality: 6},
		},
		{
			desc:     "Conjured expired item degrdaes twice as fast as normal expired",
			name:     "Conjured",
			initial:  status{sellIn: 0, quality: 8},
			expected: status{sellIn: -1, quality: 4},
		},
		{
			desc:     "Conjured expired item quality never less than 0",
			name:     "Conjured",
			initial:  status{sellIn: 0, quality: 3},
			expected: status{sellIn: -1, quality: 0},
		},
		{
			desc:     "Conjured unexpired item quality never less than 0",
			name:     "Conjured",
			initial:  status{sellIn: 2, quality: 1},
			expected: status{sellIn: 1, quality: 0},
		},
	}
	fmt.Println(len(testCases), "testcases")
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			actual := NewUpdatableItem(tC.name, tC.initial.sellIn, tC.initial.quality)
			UpdateQuality([]UpdatableItem{actual})
			expected := NewUpdatableItem(tC.name, tC.expected.sellIn, tC.expected.quality)
			if actual.Name() != expected.Name() {
				t.Errorf("expected Name to be %q, got %q", expected.Name(), actual.Name())
			}
			if actual.SellIn() != expected.SellIn() {
				t.Errorf("expected SellIn to be %d, got %d", expected.SellIn(), actual.SellIn())
			}
			if actual.Quality() != expected.Quality() {
				t.Errorf("expected Quality to be %d, got %d", expected.Quality(), actual.Quality())
			}
		})
	}
}
