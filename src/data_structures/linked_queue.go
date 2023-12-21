package datastructures

import "errors"

type LinkedQueue struct {
	start  *Node
	end    *Node
	length int
}

func (q *LinkedQueue) GetLength() int {
	return q.length
}

func (q *LinkedQueue) Enqueue(value any) error {
	newNode := &Node{value: value}

	if q.length == 0 {
		q.start = newNode
		q.end = newNode
	} else {
		q.end.next = newNode
		q.end = newNode
	}

	q.length++
	return nil
}

func (q *LinkedQueue) Dequeue() (any, error) {
	if q.length == 0 {
		return nil, errors.New("Can't dequeue empty queue!")
	}

	node := q.start

	q.start = q.start.next
	q.length--

	return node.value, nil
}

func (q *LinkedQueue) Read() (any, error) {
	if q.length == 0 {
		return nil, errors.New("Can't read from empty queue!")
	}
	return q.start.value, nil
}
