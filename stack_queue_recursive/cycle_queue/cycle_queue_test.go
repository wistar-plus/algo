package cycle_queue

import (
	"fmt"
	"testing"
)

func Test_CycleQueue(t *testing.T) {
	q := NewCycleQueue(8)
	for i := 0; i < 8; i++ {
		q.EnQueue(i)
	}
	fmt.Println(q)

	fmt.Println(q.DeQueue())
	fmt.Println(q.DeQueue())
	q.EnQueue("a")
	q.EnQueue("b")
	q.EnQueue("c")
	fmt.Println(q)
}
