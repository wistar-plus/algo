package merge_sort

//递归遍历，左边有序，右边有序，合并有序数组
func MergeSort(arr []int) {
	if len(arr) <= 1 {
		return
	}

	mergeSort(arr, 0, len(arr)-1)
}

func mergeSort(arr []int, start, end int) {
	if start >= end {
		return
	}

	mid := (start + end) / 2

	mergeSort(arr, start, mid)
	mergeSort(arr, mid+1, end)

	merge(arr, start, mid, end)
}

func merge(arr []int, start, mid, end int) {
	tmp := make([]int, end-start+1)

	i, j, k := start, mid+1, 0

	for i <= mid && j <= end {
		if arr[i] > arr[j] {
			tmp[k] = arr[j]
			j++
		} else {
			tmp[k] = arr[i]
			i++
		}
		k++
	}

	for ; i <= mid; i++ {
		tmp[k] = arr[i]
		k++
	}
	for ; j <= end; j++ {
		tmp[k] = arr[j]
		k++
	}
	copy(arr[start:end+1], tmp)
}
