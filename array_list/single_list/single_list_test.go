package single_list

import (
	"fmt"
	"testing"
)

func Test_List(t *testing.T) {
	l := NewList()
	for i := 0; i < 10; i++ {
		l.PushFront(i)
	}
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Printf("%v ", e.Value)
	}
	fmt.Println("")

	l2 := l.Reverse()
	for e := l2.Front(); e != nil; e = e.Next() {
		fmt.Printf("%v ", e.Value)
	}
	fmt.Println("")

	l2.Delete(l2.Front().Next().Next())

	for e := l2.Front(); e != nil; e = e.Next() {
		fmt.Printf("%v ", e.Value)
	}
	fmt.Println("")
	fmt.Printf("Media = %v\n", l2.Media().Value)
}
