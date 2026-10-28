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
		// TODO: add more test cases
	}
	fmt.Println(len(testCases), "testcases")
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			actual := &Item{
				Name:    tC.name,
				SellIn:  tC.initialSellIn,
				Quality: tC.initialQuality,
			}
			expected := &Item{
				Name:    tC.name,
				SellIn:  tC.expectedSellIn,
				Quality: tC.expectedQuality,
			}
			UpdateQuality([]*Item{actual})
			if actual.Name != expected.Name {
				t.Errorf("expected Name to be %q, got %q", expected.Name, actual.Name)
			}
			if actual.SellIn != expected.SellIn {
				t.Errorf("expected SellIn to be %d, got %d", expected.SellIn, actual.SellIn)
			}
			if actual.Quality != expected.Quality {
				t.Errorf("expected Quality to be %d, got %d", expected.Quality, actual.Quality)
			}
		})
	}
}
