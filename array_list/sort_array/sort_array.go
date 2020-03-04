package sort_array

import "fmt"

type array struct {
	arr []int
	len uint
	cap uint
}

func NewArray(cap uint) *array {
	return &array{arr: make([]int, 10), len: 0, cap: cap}
}

func (arr *array) Sort_Push(data int) bool {
	if arr.len >= arr.cap {
		return false
	}
	i := arr.len
	for i > 0 {
		if arr.arr[i-1] < data {
			break
		}
		arr.arr[i] = arr.arr[i-1]
		i--
	}
	arr.arr[i] = data
	arr.len++
	return true
}

func (arr *array) Print() {
	var format string
	for i := uint(0); i < arr.len; i++ {
		format += fmt.Sprintf("|%+v", arr.arr[i])
	}
	fmt.Println(format)
}

func mergeSortArray(arr1, arr2 []int) []int {
	arr := make([]int, len(arr1)+len(arr2))

	len1, len2 := len(arr1), len(arr2)

	i, j, k := 0, 0, 0
	for {
		if j == len1 {
			for k < len2 {
				arr[i] = arr2[k]
				i++
				k++
			}
			break
		}
		if k == len2 {
			for j < len1 {
				arr[i] = arr1[j]
				i++
				j++
			}
			break
		}

		if arr1[j] < arr2[k] {
			arr[i] = arr1[j]
			j++
		} else {
			arr[i] = arr2[k]
			k++
		}
		i++
	}
	return arr
}
