package prefixsum

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPrefixSum(t *testing.T) {
	c := require.New(t)

	c.Equal([]int{3, 5, 6, 11, 15}, prefixSum([]int{3, 2, 1, 5, 4}))
}

func TestRangePrefixSum(t *testing.T) {
	c := require.New(t)

	prefixSumArray := prefixSum([]int{3, 2, 1, 5, 4})

	c.Equal(10, rangePrefixSum(prefixSumArray, 2, 4))
	c.Equal(15, rangePrefixSum(prefixSumArray, 0, 4))
	c.Equal(8, rangePrefixSum(prefixSumArray, 1, 3))
}
