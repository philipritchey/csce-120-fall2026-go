package gildedrose

import (
	"fmt"
	"testing"
)

func TestUpdateQuality(t *testing.T) {
	testCases := []struct {
		desc            string
		name            string
		initialSellIn   int
		initialQuality  int
		expectedSellIn  int
		expectedQuality int
	}{
		{
			desc:            "normal item",
			name:            "normal",
			initialSellIn:   10,
			initialQuality:  20,
			expectedSellIn:  9,
			expectedQuality: 19,
		},
		{
			desc:            "normal item just before expiration",
			name:            "normal",
			initialSellIn:   1,
			initialQuality:  20,
			expectedSellIn:  0,
			expectedQuality: 19,
		},
		{
			desc:            "normal item at expiration",
			name:            "normal",
			initialSellIn:   0,
			initialQuality:  19,
			expectedSellIn:  -1,
			expectedQuality: 17,
		},
		{
			desc:            "normal item past expiration has quality not less than 0",
			name:            "normal",
			initialSellIn:   -1,
			initialQuality:  1,
			expectedSellIn:  -2,
			expectedQuality: 0,
		},
		{
			desc:            "aged brie before expiration",
			name:            "Aged Brie",
			initialSellIn:   10,
			initialQuality:  20,
			expectedSellIn:  9,
			expectedQuality: 21,
		},
		{
			desc:            "aged brie before expiration has quality not more than 50",
			name:            "Aged Brie",
			initialSellIn:   10,
			initialQuality:  50,
			expectedSellIn:  9,
			expectedQuality: 50,
		},
		{
			desc:            "aged brie after expiration improves twice as fast",
			name:            "Aged Brie",
			initialSellIn:   0,
			initialQuality:  30,
			expectedSellIn:  -1,
			expectedQuality: 32,
		},
		{
			desc:            "aged brie just before expiration improves like normal",
			name:            "Aged Brie",
			initialSellIn:   1,
			initialQuality:  30,
			expectedSellIn:  0,
			expectedQuality: 31,
		},
		{
			desc:            "aged brie after expiration has quality not more than 50",
			name:            "Aged Brie",
			initialSellIn:   -2,
			initialQuality:  50,
			expectedSellIn:  -3,
			expectedQuality: 50,
		},
		{
			desc:            "Sulfuras never has to be sold and never degrades in quality",
			name:            "Sulfuras, Hand of Ragnaros",
			initialSellIn:   10,
			initialQuality:  80,
			expectedSellIn:  10,
			expectedQuality: 80,
		},
		{
			desc:            "Sulfuras after expiration never has to be sold and never degrades in quality",
			name:            "Sulfuras, Hand of Ragnaros",
			initialSellIn:   -1,
			initialQuality:  80,
			expectedSellIn:  -1,
			expectedQuality: 80,
		},
		{
			desc:            "Backstage passes with more than 10 days improve by 1",
			name:            "Backstage passes to a TAFKAL80ETC concert",
			initialSellIn:   11,
			initialQuality:  30,
			expectedSellIn:  10,
			expectedQuality: 31,
		},
		{
			desc:            "Backstage passes with 10 days improve by 2",
			name:            "Backstage passes to a TAFKAL80ETC concert",
			initialSellIn:   10,
			initialQuality:  30,
			expectedSellIn:  9,
			expectedQuality: 32,
		},
		{
			desc:            "Backstage passes with 6 days improve by 2",
			name:            "Backstage passes to a TAFKAL80ETC concert",
			initialSellIn:   6,
			initialQuality:  30,
			expectedSellIn:  5,
			expectedQuality: 32,
		},
		{
			desc:            "Backstage passes with 5 days improve by 3",
			name:            "Backstage passes to a TAFKAL80ETC concert",
			initialSellIn:   5,
			initialQuality:  30,
			expectedSellIn:  4,
			expectedQuality: 33,
		},
		{
			desc:            "Backstage passes with 1 day improve by 3",
			name:            "Backstage passes to a TAFKAL80ETC concert",
			initialSellIn:   1,
			initialQuality:  30,
			expectedSellIn:  0,
			expectedQuality: 33,
		},
		{
			desc:            "Backstage passes expired have quality 0",
			name:            "Backstage passes to a TAFKAL80ETC concert",
			initialSellIn:   0,
			initialQuality:  30,
			expectedSellIn:  -1,
			expectedQuality: 0,
		},
		{
			desc:            "Backstage passes quality never less than 0",
			name:            "Backstage passes to a TAFKAL80ETC concert",
			initialSellIn:   -1,
			initialQuality:  0,
			expectedSellIn:  -2,
			expectedQuality: 0,
		},
		{
			desc:            "Backstage passes quality never more than 50 when far away",
			name:            "Backstage passes to a TAFKAL80ETC concert",
			initialSellIn:   20,
			initialQuality:  50,
			expectedSellIn:  19,
			expectedQuality: 50,
		},
		{
			desc:            "Backstage passes quality never more than 50 when near",
			name:            "Backstage passes to a TAFKAL80ETC concert",
			initialSellIn:   8,
			initialQuality:  50,
			expectedSellIn:  7,
			expectedQuality: 50,
		},
		{
			desc:            "Backstage passes quality never more than 50 when very soon",
			name:            "Backstage passes to a TAFKAL80ETC concert",
			initialSellIn:   3,
			initialQuality:  50,
			expectedSellIn:  2,
			expectedQuality: 50,
		},
		{
			desc:            "Conjured unexpired degrades twice as fast as normal unexpired",
			name:            "Conjured",
			initialSellIn:   10,
			initialQuality:  20,
			expectedSellIn:  9,
			expectedQuality: 18,
		},
		{
			desc:            "Conjured expired degrades twice as fast as normal expired",
			name:            "Conjured",
			initialSellIn:   0,
			initialQuality:  20,
			expectedSellIn:  -1,
			expectedQuality: 16,
		},
		{
			desc:            "Conjured unexpired quality not less than 0",
			name:            "Conjured",
			initialSellIn:   10,
			initialQuality:  1,
			expectedSellIn:  9,
			expectedQuality: 0,
		},
		{
			desc:            "Conjured expired quality not less than 0",
			name:            "Conjured",
			initialSellIn:   0,
			initialQuality:  3,
			expectedSellIn:  -1,
			expectedQuality: 0,
		},
	}

	fmt.Println(len(testCases), "testcases")

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			// Given an item
			actual := NewUpdatableItem(tC.name, tC.initialSellIn, tC.initialQuality)

			// When I update quality
			UpdateQuality([]UpdatableItem{actual})

			// Then the item's fields should be updated correctly
			expected := NewUpdatableItem(tC.name, tC.expectedSellIn, tC.expectedQuality)
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
