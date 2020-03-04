package cycle_list

import (
	"fmt"
	"testing"
)

func Test_CycleList(t *testing.T) {
	arr := NewCycleList()
	for i := 0; i < 5; i++ {
		arr.PushFront(i + 1)
	}

	e := arr.Front()
	for i := 0; i < 15; i++ {
		fmt.Printf("%v ", e.Value)
		e = e.Next()
	}
	fmt.Println()

	arr.Delete(arr.Front())
	e = arr.Front()
	for i := 0; i < 15; i++ {
		fmt.Printf("%v ", e.Value)
		e = e.Next()
	}
	fmt.Println()

	fmt.Println(arr.hasLoop())

}
