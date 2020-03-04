package dynamic_array

import (
	"testing"
)

func Test_DynamicArray(t *testing.T) {
	arr := NewArray(10)

	for i := 0; i < 10; i++ {
		arr.Push(i)
	}
	arr.Print()

	arr.Push("hello")
	arr.Print()

	for i := 0; i < 6; i++ {
		arr.Delete(0)
	}
	arr.Print()

}
