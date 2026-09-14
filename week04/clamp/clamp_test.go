package main

import "testing"

func TestClamp(t *testing.T) {
	type TestCase struct {
		in  int
		out int
	}
	tests := []struct {
		r     Range
		tests []TestCase
	}{
		{Range{0, 100}, []TestCase{{67, 67}, {112, 100}, {-42, 0}}},
		{Range{-10, 20}, []TestCase{{-11, -10}, {7, 7}, {31, 20}}},
	}
	for _, testset := range tests {
		r := testset.r
		for _, test := range testset.tests {
			actual := clamp(r, test.in)
			expected := test.out
			if actual != expected {
				t.Errorf("expected clamp(%v, %d) to be %d, got %d", r, test.in, expected, actual)
			}
		}
	}
}
