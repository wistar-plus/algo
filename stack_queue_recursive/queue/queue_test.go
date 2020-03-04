package queue

import (
	"fmt"
	"testing"
)

func Test_Queue(t *testing.T) {
	q := NewQueue(10)

	for i := 0; i < 10; i++ {
		q.EnQueue(i)
	}

	fmt.Println(q.EnQueue(10))

	fmt.Println(q.DeQueue())

	fmt.Println(q.EnQueue(10))

	for i := 0; i < 10; i++ {
		fmt.Printf("%v ", q.DeQueue())
	}
	fmt.Printf("%v ", q.DeQueue())
}
