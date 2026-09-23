package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCombinationLock(t *testing.T) {
	tests := []struct {
		initial  int
		first    int
		second   int
		third    int
		expected int
	}{
		{0, 30, 0, 30, 1350},
		{20, 26, 9, 23, 1827},
	}

	for _, test := range tests {
		require.Equal(t, test.expected, CombinationLock(test.initial, test.first, test.second, test.third))
	}
}
