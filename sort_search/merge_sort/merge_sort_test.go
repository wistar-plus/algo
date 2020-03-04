package merge_sort

import (
	"fmt"
	"testing"
)

func Test_MergeSort(t *testing.T) {
	arr := []int{5, 6, 8, 7, 9, 1, 3, 2, 4, 0}
	fmt.Println(arr)
	MergeSort(arr)
	fmt.Println(arr)
}
