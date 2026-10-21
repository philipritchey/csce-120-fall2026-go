package sumsquaredifference

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSumSquaredDifferece(t *testing.T) {
	assert.Equal(t, SumOfSquares(10), 385)
	assert.Equal(t, SquareOfSum(10), 3025)
	require.Equal(t, SumSquareDifference(10), 2640)

	assert.Equal(t, SumOfSquares(100), 338_350)
	assert.Equal(t, SquareOfSum(100), 25_502_500)
	require.Equal(t, SumSquareDifference(100), 25_164_150)
}
