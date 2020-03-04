package selection_sort

//两次for循环，遍历未排序区间，选出最大或者最小放到已排序区间
func SelectionSort(arr []int) {
	if len(arr) <= 1 {
		return
	}
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
}
