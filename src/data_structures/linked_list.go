package datastructures

import (
	"errors"
)

type LinkedList struct {
	head   *Node
	length int
}

func (l *LinkedList) GetLength() int {
	return l.length
}

func (l *LinkedList) Insert(value any, position int) error {
	if position < 0 || position > l.length {
		return errors.New("Invalid position!")
	}

	if l.head == nil || position == 0 {
		l.head = &Node{value: value, next: l.head}
	} else {
		currentNode := l.head
		for i := 0; i < position-1; i++ {
			currentNode = currentNode.next
		}
		currentNode.next = &Node{
			value: value,
			next:  currentNode.next,
		}
	}

	l.length++
	return nil
}

func (l *LinkedList) Delete(position int) error {
	if position < 0 || position >= l.length {
		return errors.New("Invalid position!")
	}

	if position == 0 {
		l.head = l.head.next
	} else {
		currentNode := l.head
		for i := 0; i < position-1; i++ {
			currentNode = currentNode.next
		}
		currentNode.next = currentNode.next.next
	}

	l.length--
	return nil
}

func (l *LinkedList) Consult(position int) (any, error) {
	if position < 0 || position >= l.length {
		return nil, errors.New("Invalid position!")
	}

	currentNode := l.head
	for i := 0; i < position; i++ {
		currentNode = currentNode.next
	}

	return currentNode.value, nil
}

func (l *LinkedList) Search(value any) (int, error) {
	currentNode := l.head
	for i := 0; currentNode != nil; i++ {
		if currentNode.value == value {
			return i, nil
		}
		currentNode = currentNode.next
	}

	return -1, errors.New("Value not found!")
}
