package prefixsum

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMainFunction(t *testing.T) {
	c := require.New(t)

	c.Equal([]int{3, 5, 6, 11, 15}, prefixSum([]int{3, 2, 1, 5, 4}))
}
