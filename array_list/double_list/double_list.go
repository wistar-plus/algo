package double_list

type doubleList struct {
	head   *ListNode
	tail   *ListNode
	length int
}

type ListNode struct {
	Value interface{}
	next  *ListNode
	prev  *ListNode
}

func NewListNode(value interface{}) *ListNode {
	return &ListNode{Value: value, next: nil}
}

func NewList() *doubleList {
	return &doubleList{head: NewListNode(0), length: 0, tail: NewListNode(0)}
}

func (l *doubleList) PushFront(value interface{}) {
	t := l.head.next
	n := NewListNode(value)
	n.next = t
	if t != nil {
		t.prev = n
	}
	l.head.next = n
	if l.length == 0 {
		l.tail.prev = n
	}
	l.length++
}

func (l *doubleList) Front() *ListNode {
	return l.head.next
}

func (l *doubleList) Back() *ListNode {
	return l.tail.prev
}

func (l *ListNode) Next() *ListNode {
	return l.next
}

func (l *ListNode) Prev() *ListNode {
	return l.prev
}

func (l *doubleList) Delete(n *ListNode) bool {
	if n == nil {
		return false
	}
	for e := l.head; e.next != nil; e = e.Next() {
		if e.next == n {
			e.next = n.next
			if n.next != nil {
				n.next.prev = n.prev
			} else {
				l.tail.prev = n.prev
			}
			l.length--
			return true
		}
	}
	return false
}
