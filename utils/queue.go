package utils

type queueNode[T any] struct {
	data T
	next *queueNode[T]
}

type Queue[T any] struct {
	head  *queueNode[T]
	tail  *queueNode[T]
	count int
}

func (que *Queue[T]) Len() int {
	return que.count
}

func (que *Queue[T]) Push(item T) {
	next := &queueNode[T]{data: item}
	if que.head == nil {
		que.tail = next
		que.head = next
	} else {
		que.tail.next = next
		que.tail = next
	}
	que.count++
}

func (que *Queue[T]) Pop() T {

	// if que.head == nil {
	// 	return nil
	// }

	front := que.head
	que.head = front.next

	if que.head == nil {
		que.tail = nil
	}
	que.count--

	return front.data
}
