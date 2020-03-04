package single_list

type List struct {
	head   *ListNode
	length int
}

type ListNode struct {
	Value interface{}
	next  *ListNode
}

func NewListNode(value interface{}) *ListNode {
	return &ListNode{Value: value, next: nil}
}

func NewList() *List {
	return &List{head: NewListNode(0), length: 0}
}

func (l *List) PushFront(value interface{}) {
	t := l.head.next
	l.head.next = NewListNode(value)
	l.head.next.next = t
	l.length++
}

func (l *List) Front() *ListNode {
	return l.head.next
}

func (l *ListNode) Next() *ListNode {
	return l.next
}

func (l *List) Delete(n *ListNode) bool {
	if n == nil {
		return false
	}
	for e := l.head; e.next != nil; e = e.Next() {
		if e.next == n {
			e.next = n.next
			l.length--
			return true
		}
	}
	return false
}

func (l *List) Reverse() *List {
	t := NewList()
	for e := l.Front(); e != nil; e = e.Next() {
		t.PushFront(e.Value)
	}
	return t
}

func (l *List) Media() *ListNode {
	q, s := l.head.next, l.head.next
	for q != nil && q.next != nil {
		s = s.next
		q = q.next.next
	}
	return s
}
