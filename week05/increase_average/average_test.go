package increaseaverage

import (
	"bytes"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAverage(t *testing.T) {
	tests := []struct {
		a   []int
		b   []int
		out int
	}{
		{[]int{100, 101, 102, 103, 104}, []int{98, 100, 102, 99, 101}, 1},
		{[]int{42, 67, 98}, []int{12}, 2},
		{[]int{86, 75, 39}, []int{67}, 0},
	}
	for _, test := range tests {
		actual := averageIncreaseCount(test.a, test.b)
		expected := test.out
		require.Equal(t, expected, actual)
	}
}

func TestMainMethod(t *testing.T) {
	reader := strings.NewReader("3 1\n42 67 98\n12")
	var writer bytes.Buffer
	Main(reader, &writer)
	require.Equal(t, "2", strings.TrimSpace(writer.String()))
}
