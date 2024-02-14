package quick_sort

func quickSort(arr []int) {

}

func partition(arr []int, low, hi int) int {
	pivot := arr[len(arr)/2]
	idx := -1
	tmp := 0

	for i := low; i < hi; i++ {
		if arr[i] <= pivot {
			idx++
			tmp = arr[idx]
			arr[idx] = arr[i]
			arr[i] = tmp
		}
	}

	idx++
	tmp = arr[idx]
	arr[idx] = pivot
	arr[len(arr)/2] = tmp

	return 0
}

func qs(arr []int, low, hi int) {

}

func getPivot(arr []int) int {
	return len(arr) / 2
}
