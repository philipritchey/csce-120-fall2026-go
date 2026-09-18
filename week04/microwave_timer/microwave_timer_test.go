package main

import "testing"

func TestMicrowaveTimer(t *testing.T) {
	tests := []struct {
		name string
		in   int
		out  int
	}{
		{"0:32", 32, 32}, // n < 100 -> 00:ss -> n = ss -> n
		{"0:67", 60, 60},
		{"0:67", 67, 67},
		{"0:83", 83, 83},
		{"0:99", 99, 99},
		{"1:00", 100, 60}, // n >= 100 -> mm:ss -> n = (mm)(ss) -> 60*(mm)+(ss)
		{"1:07", 107, 67},
		{"1:23", 123, 83},
		{"0:99", 139, 99},
		{"5:67", 567, 367},
		{"9:99", 999, 639},
		{"12:34", 1234, 754},
		{"59:59", 5959, 3599},
		{"60:00", 6000, 3600},
		{"99:99", 9999, 6039},
	}
	for _, test := range tests {
		actual := MicrowaveTimer(test.in)
		expected := test.out
		if actual != expected {
			t.Errorf("%s: expected %d, got %d", test.name, expected, actual)
		}
	}
}
