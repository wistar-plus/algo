package stack

import (
	"fmt"
	"testing"
)

func Test_ArrayStack(t *testing.T) {
	s := NewArrayStack(10)
	for i := 0; i < 10; i++ {
		s.Push(i)
	}
	fmt.Println(s)
	for i := 0; i < 10; i++ {
		fmt.Printf("%v ", s.Pop())
	}
	fmt.Println()
	fmt.Println(s)
}

func Test_ListStack(t *testing.T) {
	s := NewListStack(10)
	for i := 0; i < 10; i++ {
		s.Push(i)
	}
	for i := 0; i < 10; i++ {
		fmt.Printf("%v ", s.Pop())
	}
}
