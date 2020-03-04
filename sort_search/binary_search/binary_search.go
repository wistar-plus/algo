package binary_search

func BinarySearch(arr []int, val int) int {

	low, high := 0, len(arr)-1

	for low <= high {
		mid := low + ((high - low) >> 1)
		if arr[mid] == val {
			if mid == 0 || arr[mid-1] != val {
				return mid
			} else {
				high = mid - 1
			}
		} else if arr[mid] < val {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}
