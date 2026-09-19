package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAbc(t *testing.T) {
	slice := []int{8, 6, 7, 5, 3, 0, 9}
	originalSlice := make([]int, len(slice))
	copy(originalSlice, slice)
	a, b, c := abc(len(slice), slice)
	require.NotEqual(t, a, b, "expected distinct indexes, got index a == index b")
	require.NotEqual(t, a, c, "expected distinct indexes, got index a == index c")
	require.NotEqual(t, b, c, "expected distinct indexes, got index b == index c")
	require.Equal(t, originalSlice[c], originalSlice[a]+originalSlice[b], "%d + %d != %d", originalSlice[a], originalSlice[b], originalSlice[c])
	require.Equal(t, originalSlice, slice, "slice was modified")

	slice = []int{1, 2, 4, 7, 10, 13, 16, 19, 22}
	originalSlice = make([]int, len(slice))
	copy(originalSlice, slice)
	a, _, _ = abc(len(slice), slice)
	require.Equal(t, -1, a, "expected index a to be -1, got %d", a)

	slice = []int{1, 2}
	originalSlice = make([]int, len(slice))
	copy(originalSlice, slice)
	a, _, _ = abc(len(slice), slice)
	require.Equal(t, -1, a, "expected index a to be -1, got %d", a)
}

func TestMainImposible(t *testing.T) {
	reader := strings.NewReader("2\n1 2")
	var writer strings.Builder
	Main(reader, &writer)
	require.Equal(t, "impossible", strings.TrimSpace(writer.String()))
}

func TestMainSolutionFound(t *testing.T) {
	reader := strings.NewReader("3\n1 2 3")
	var writer strings.Builder
	Main(reader, &writer)
	require.Equal(t, "1 + 2 = 3\n0, 1, 2", strings.TrimSpace(writer.String()))
}

func TestMainIgnoreExtraInput(t *testing.T) {
	testcases := []struct {
		in  string
		out string
	}{
		{"3\n1 2 3 4 5", "1 + 2 = 3\n0, 1, 2"},
		{"3\n5 4 3 2 1", "impossible"},
	}
	for _, test := range testcases {
		reader := strings.NewReader(test.in)
		var writer strings.Builder
		Main(reader, &writer)
		require.Equal(t, test.out, strings.TrimSpace(writer.String()), "given %q, expected %q", test.in, test.out)
	}
}

func TestMainInvalidInput(t *testing.T) {
	testcases := []string{
		"3\n1 2 x 3",
		"\n",
		"",
		"-1",
	}
	for _, input := range testcases {
		reader := strings.NewReader(input)
		var writer strings.Builder
		Main(reader, &writer)
		require.Contains(t, writer.String(), "invalid input", "expected %q to be invalid", input)
	}
}

func TestMainMultilineInput(t *testing.T) {
	reader := strings.NewReader("3\n1 2\n3")
	var writer strings.Builder
	Main(reader, &writer)
	require.Equal(t, "1 + 2 = 3\n0, 1, 2", strings.TrimSpace(writer.String()))
}

func TestPrintOutput(t *testing.T) {
	var writer strings.Builder
	printOutput(&writer, []int{1, 2, 3}, 0, 1, 2)
	require.Equal(t, "1 + 2 = 3\n0, 1, 2", strings.TrimSpace(writer.String()))
}

func TestGetInput(t *testing.T) {
	reader := strings.NewReader("3\n1 2 3")
	n, slice, err := getInput(reader)
	require.Nil(t, err)
	require.Equal(t, 3, n)
	require.Equal(t, []int{1, 2, 3}, slice)
}
