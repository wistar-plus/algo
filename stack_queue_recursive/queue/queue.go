package queue

type Queue struct {
	arr  []interface{}
	head int
	tail int
	cap  int
}

func NewQueue(cap int) *Queue {
	return &Queue{arr: make([]interface{}, cap), cap: cap}
}

func (q *Queue) EnQueue(data interface{}) bool {
	if q.tail == q.cap {
		if q.head == 0 {
			return false
		}
		for i := q.head; i < q.tail; i++ {
			q.arr[i-q.head] = q.arr[i]
		}
		q.tail -= q.head
		q.head = 0
	}
	q.arr[q.tail] = data
	q.tail++
	return true
}

func (q *Queue) DeQueue() interface{} {
	if q.head < q.tail {
		ret := q.arr[q.head]
		q.head++
		return ret
	}
	return nil
}

// type ListQueue struct {

// }
