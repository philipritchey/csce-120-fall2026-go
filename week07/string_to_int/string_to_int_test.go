package stringtoint

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStringToIntValid(t *testing.T) {
	tests := []struct {
		in  string
		out int
	}{
		{"0", 0},
		{"1", 1},
		{"8675309", 8675309},
		{"-979", -979},
		{"0099", 99},
	}
	for _, tc := range tests {
		actual, err := StringToInt(tc.in)
		require.Nil(t, err, "%q should be valid", tc.in)
		require.Equal(t, tc.out, actual, "got %d, want %d", actual, tc.out)
	}
}

func TestStringToIntInvalid(t *testing.T) {
	invalidInputs := []string{
		"",
		"--979",
		"x",
		"12x4",
		"3.14",
		"9223372036854775808",
	}
	for _, tc := range invalidInputs {
		_, err := StringToInt(tc)
		require.NotNil(t, err, "%q should be invalid", tc)
	}
}
