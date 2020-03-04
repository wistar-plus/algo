package bubble_sort

//两次for循环，比较相邻两个数，看是否符合大小要求，不符合就互换位置
func BubbleSort(arr []int) {
	if len(arr) <= 1 {
		return
	}
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}
