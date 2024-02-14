package quick_sort

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestQuickSort(t *testing.T) {
	c := require.New(t)

	arr := []int{5, 6, 3, 8, 1, 2, 7, 9, 4}
	quickSort(arr)
	c.Equal([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, arr)
}
