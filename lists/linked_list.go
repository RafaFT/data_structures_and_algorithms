package lists

import (
	"fmt"
	"iter"
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

	for i, v := range l.all() {
		if i == index {
			return v.value
		}
	}

	panic("unreachable")
}

// Search returns the first index that contains value or -1.
//
// Time O(n) and space O(1).
func (l *LinkedList[T]) Search(value T) int {
	for i, v := range l.all() {
		if v.value == value {
			return i
		}
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

	var prevNode *node[T]
	for i, v := range l.all() {
		if i == index-1 {
			prevNode = v
		}
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

	var prevNode *node[T]
	for i, v := range l.all() {
		if i == index-1 {
			prevNode = v
		}
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

// all returns an iterator over LinkedList's index-node pairs.
func (l *LinkedList[T]) all() iter.Seq2[int, *node[T]] {
	return func(yield func(int, *node[T]) bool) {
		i := 0
		for currentNode := l.head; currentNode != nil; currentNode = currentNode.next {
			if !yield(i, currentNode) {
				return
			}
			i++
		}
	}
}

// All returns an iterator over LinkedList's index-value pairs.
func (l *LinkedList[T]) All() iter.Seq2[int, T] {
	return func(yield func(int, T) bool) {
		for i, v := range l.all() {
			if !yield(i, v.value) {
				return
			}
		}
	}
}

// Values returns an iterator over LinkedList's elements.
func (l *LinkedList[T]) Values() iter.Seq[T] {
	return func(yield func(T) bool) {
		for _, value := range l.All() {
			if !yield(value) {
				return
			}
		}
	}
}
