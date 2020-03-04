package quick_sort

//递归遍历，小于分区点放左边，大于放右边
func QuickSort(arr []int) {
	if len(arr) <= 1 {
		return
	}
	quickSort(arr, 0, len(arr)-1)
}

func quickSort(arr []int, start, end int) {
	if start >= end {
		return
	}

	pivot := arr[end]
	idx := start

	for i := start; i < end; i++ {
		if arr[i] < pivot {
			arr[idx], arr[i] = arr[i], arr[idx]
			idx++
		}
	}
	arr[idx], arr[end] = arr[end], arr[idx]

	quickSort(arr, start, idx-1)
	quickSort(arr, idx+1, end)
}
