package cycle_queue

type CycleQueue struct {
	arr  []interface{}
	head int
	tail int
	cap  int
}

func NewCycleQueue(cap int) *CycleQueue {
	return &CycleQueue{arr: make([]interface{}, cap+1), cap: cap + 1}
}

func (q *CycleQueue) EnQueue(data interface{}) bool {
	if (q.tail+1)%q.cap == q.head {
		return false
	}
	q.arr[q.tail] = data
	q.tail = (q.tail + 1) % q.cap
	return true
}

func (q *CycleQueue) DeQueue() interface{} {
	if q.head == q.tail {
		return nil
	}
	ret := q.arr[q.head]
	q.head = (q.head + 1) % q.cap
	return ret
}
