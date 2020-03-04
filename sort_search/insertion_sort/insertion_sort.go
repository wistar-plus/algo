package insertion_sort

//两次for循环，从未排序区间取出一个元素插入到已排序区间
func InsertionSort(arr []int) {
	if len(arr) <= 1 {
		return
	}
	for i := 1; i < len(arr); i++ {
		tmp := arr[i]
		j := i - 1
		for j >= 0 {
			if tmp < arr[j] {
				arr[j+1] = arr[j]
			} else {
				break
			}
			j--
		}
		arr[j+1] = tmp
	}
}
