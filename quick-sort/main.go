package quick_sort

func quickSort(arr []int) {
	qs(arr, 0, len(arr)-1)
	return
}

func qs(arr []int, low, hi int) {
	if low >= hi {
		return
	}

	partitionIdx := partition(arr, low, hi)

	qs(arr, low, partitionIdx-1)
	qs(arr, partitionIdx+1, hi)
}

func partition(arr []int, low, hi int) int {
	pivotIdx := hi
	pivot := arr[pivotIdx]
	idx := low - 1
	tmp := 0

	for i := low; i < hi; i++ {
		if arr[i] <= pivot {
			idx++
			tmp = arr[i]
			arr[i] = arr[idx]
			arr[idx] = tmp
		}
	}

	idx++
	arr[pivotIdx] = arr[idx]
	arr[idx] = pivot

	return idx
}
