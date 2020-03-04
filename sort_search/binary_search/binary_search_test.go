package binary_search

import (
	"fmt"
	"testing"
)

func Test_BinarySearch(t *testing.T) {
	arr := []int{0, 1, 2, 3, 4, 4, 5, 6, 7, 8, 9}

	fmt.Println(BinarySearch(arr, 0))
	fmt.Println(BinarySearch(arr, 4))
	fmt.Println(BinarySearch(arr, 20))
}
