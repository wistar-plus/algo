package sort_array

import (
	"fmt"
	"testing"
)

func Test_SortArray(t *testing.T) {
	arr := NewArray(10)
	tmp := []int{4, 2, 6, 7, 1, 9, 5, 8, 3, 0}
	for _, v := range tmp {
		arr.Sort_Push(v)
		arr.Print()
	}
}

func Test_MegerSortArray(t *testing.T) {
	arr1 := []int{1, 3, 5, 7, 8, 9}
	arr2 := []int{0, 2, 4, 6}

	arr := mergeSortArray(arr1, arr2)

	fmt.Println(arr)
}
