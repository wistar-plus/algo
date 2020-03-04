package double_list

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

	l.Delete(l.Front())
	l.Delete(l.Back())
	for e := l.Back(); e != nil; e = e.Prev() {
		fmt.Printf("%v ", e.Value)
	}
	fmt.Println("")
}
