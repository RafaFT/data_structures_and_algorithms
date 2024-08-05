package lists

import (
	"fmt"
	"strings"
)

type LinkedList[T comparable] struct {
	len  int
	head *node[T]
}

type node[T comparable] struct {
	value T
	next  *node[T]
}

// Read returns the value at the provided index.
// It panic if index is out of bounds.
//
// Time O(n) and space O(1).
func (l *LinkedList[T]) Read(index int) T {
	if index < 0 || index >= l.len {
		panic(fmt.Sprintf("index out of range [%d] with length %d", index, l.len))
	}

	currentNode := l.head
	for i := 1; i <= index; i++ {
		currentNode = currentNode.next
	}

	return currentNode.value
}

// Search returns the first index that contains value or -1.
//
// Time O(n) and space O(1).
func (l *LinkedList[T]) Search(value T) int {
	i := 0
	for currentNode := l.head; currentNode != nil; currentNode = currentNode.next {
		if currentNode.value == value {
			return i
		}
		i++
	}

	return -1
}

// Insert inserts value at the provided index.
// It panics if index is out of range.
//
// Time O(n) and space O(1).
func (l *LinkedList[T]) Insert(value T, index int) {
	if index < 0 || index > l.len {
		panic(fmt.Sprintf("index out of range [%d] with length %d", index, l.len))
	}

	l.len++
	newNode := &node[T]{
		value: value,
		next:  nil,
	}

	if index == 0 {
		newNode.next = l.head
		l.head = newNode
		return
	}

	prevNode := l.head
	for i := 1; i < index; i++ {
		prevNode = prevNode.next
	}

	newNode.next = prevNode.next
	prevNode.next = newNode
}

// Delete removes value at provided index.
// It panics if index is out of bounds.
//
// Time O(n) and space O(1).
func (l *LinkedList[T]) Delete(index int) {
	if index < 0 || index >= l.len {
		panic(fmt.Sprintf("index out of range [%d] with length %d", index, l.len))
	}

	l.len--
	if l.len == 0 {
		l.head = nil
		return
	}

	if index == 0 {
		l.head = l.head.next
		return
	}

	prevNode := l.head
	for i := 1; i < index; i++ {
		prevNode = prevNode.next
	}

	prevNode.next = prevNode.next.next
}

func (l *LinkedList[T]) String() string {
	var builder strings.Builder

	// assume each element requires at least one byte for printing
	// and one byte for spacing between elements
	// len("LinkedList[]") + (LinkedList.len * 2)
	builder.Grow(12 + (l.len * 2))

	builder.WriteString("LinkedList[")

	i := 0
	for currentNode := l.head; currentNode != nil; currentNode = currentNode.next {
		if i != 0 {
			builder.WriteString(" ")
		}
		builder.WriteString(fmt.Sprint(currentNode.value))
		i++
	}

	builder.WriteString("]")

	return builder.String()
}
