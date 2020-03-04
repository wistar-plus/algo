package insertion_sort

import (
	"fmt"
	"testing"
)

func Test_InsertionSort(t *testing.T) {
	arr := []int{5, 6, 8, 7, 9, 1, 3, 2, 4, 0}
	fmt.Println(arr)
	InsertionSort(arr)
	fmt.Println(arr)
}
