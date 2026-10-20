package ssd

import (
	"fmt"
	"testing"
)

func TestSsd(t *testing.T) {
	testcases := []struct {
		b        int
		n        int
		expected int
	}{
		// examples from spec
		{10, 1234, 30},
		{3, 98765, 19},
		{16, 987654321, 696},
		// edge cases
		{2, 13, 3}, // 13_10 = 1101_2
		{7, 0, 0},  // SSD(b, 0) = 0 for all bases
	}
	for _, test := range testcases {
		name := fmt.Sprintf("SSD(%d,%d)", test.b, test.n)
		t.Run(name, func(t *testing.T) {
			actual := Ssd(test.b, test.n)
			if actual != test.expected {
				t.Errorf("got %d, wanted %d", actual, test.expected)
			}
		})
	}
}
