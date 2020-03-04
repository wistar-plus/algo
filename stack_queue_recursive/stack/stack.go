package stack

import "container/list"

type ArrayStack struct {
	arr []interface{}
	top int
	cap int
}

type ListStack struct {
	arr *list.List
	top int
	cap int
}

func NewArrayStack(cap int) *ArrayStack {
	return &ArrayStack{arr: make([]interface{}, cap), cap: cap}
}

func NewListStack(cap int) *ListStack {
	return &ListStack{arr: list.New(), cap: cap}
}

func (s *ArrayStack) Push(data interface{}) {
	if s.top >= s.cap {
		return
	}
	s.arr[s.top] = data
	s.top++
}

func (s *ArrayStack) Pop() interface{} {
	if s.top <= 0 {
		return nil
	}
	s.top--
	return s.arr[s.top]
}

func (s *ListStack) Push(data interface{}) {
	if s.top >= s.cap {
		return
	}
	s.arr.PushFront(data)
	s.top++
}

func (s *ListStack) Pop() interface{} {
	if s.top <= 0 {
		return nil
	}
	s.top--
	return s.arr.Remove(s.arr.Front())
}
