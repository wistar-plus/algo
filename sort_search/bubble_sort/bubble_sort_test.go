package bubble_sort

import (
	"fmt"
	"testing"
)

func Test_BubbleSort(t *testing.T) {
	arr := []int{5, 6, 8, 7, 9, 1, 3, 2, 4, 0}
	fmt.Println(arr)
	BubbleSort(arr)
	fmt.Println(arr)
}
