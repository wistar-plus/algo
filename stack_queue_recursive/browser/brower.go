package browser

import "go-learn/internal/algo/stack_queue_recursive/stack"

type Browser struct {
	prev *stack.ArrayStack
	next *stack.ArrayStack
}

func NewBrowser() *Browser {
	return &Browser{prev: stack.NewArrayStack(10), next: stack.NewArrayStack(10)}
}

func (b *Browser) View(data interface{}) {
	b.prev.Push(data)
}

func (b *Browser) Prev() interface{} {
	if res := b.prev.Pop(); res != nil {
		b.next.Push(res)
		return res
	}
	return nil
}

func (b *Browser) Next() interface{} {

	if res := b.next.Pop(); res != nil {
		b.prev.Push(res)
		return res
	}
	return nil
}
