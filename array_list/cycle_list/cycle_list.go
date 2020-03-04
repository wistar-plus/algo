package cycle_list

type cycleList struct {
	head   *ListNode
	tail   *ListNode
	length int
}

type ListNode struct {
	Value interface{}
	next  *ListNode
}

func NewListNode(value interface{}) *ListNode {
	return &ListNode{Value: value, next: nil}
}

func NewCycleList() *cycleList {
	return &cycleList{head: NewListNode(0), length: 0}
}

func (l *cycleList) PushFront(value interface{}) {
	t := l.head.next
	l.head.next = NewListNode(value)
	l.head.next.next = t
	if l.length == 0 {
		l.tail = l.head.next
	}
	l.tail.next = l.head.next
	l.length++
}

func (l *cycleList) Front() *ListNode {
	return l.head.next
}

func (l *ListNode) Next() *ListNode {
	return l.next
}

func (l *cycleList) Delete(n *ListNode) bool {
	if n == nil {
		return false
	}
	e := l.head
	for i := 0; i < l.length; i++ {
		if e.next == n {
			e.next = n.next
			if i == 0 {
				l.tail.next = n.next
			}
			l.length--
			return true
		}
		e = e.next
	}
	return false
}

func (l *cycleList) hasLoop() bool {
	q, s := l.head.next, l.head.next
	for q != nil && q.next != nil {
		q = q.next.next
		s = s.next
		if s == q {
			return true
		}
	}
	return false
}
