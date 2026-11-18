package countmatches

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func CountMatches[T comparable](values []T, target T) int {
	// return number of times target appears
	count := 0
	for _, value := range values {
		if value == target {
			count++
		}
	}
	return 0
}

// another generic test function
func AssertCountMatchesEquals[T comparable](t *testing.T, values []T, target T, expected int) {
	assert.Equal(t, expected, CountMatches(values, target), "CountMatches(%v, %v)", values, target)
}

func TestCountMatches(t *testing.T) {
	AssertCountMatchesEquals(t, []int{1, 2, 2, 3, 2}, 2, 3)
	AssertCountMatchesEquals(t, []string{"a", "b", "a"}, "a", 2)
}
