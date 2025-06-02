package arrays

import (
	"cmp"
	"fmt"
	"iter"

	"dsa/bisect"
)

// OrderedArray is an array implementation that guarantees it's values
// are kept in ascending order.
type OrderedArray[T cmp.Ordered] struct {
	arr []T
}

// Read returns the value at the provided index.
// It panics if index < 0 or index >= len(a).
//
// Time O(1) and space O(1).
func (a *OrderedArray[T]) Read(index int) T {
	if index < 0 || index >= len(a.arr) {
		panic(fmt.Sprintf("index out of range [%d] with length %d", index, len(a.arr)))
	}

	return a.arr[index]
}

// Search returns the first index that contains value or -1.
//
// Time O(log(n)) and space O(1).
func (a *OrderedArray[T]) Search(value T) int {
	i := bisect.BisectLeft(a.arr, value)

	// Check if i is within bounds and if the element at i is indeed the value we are searching for.
	if i < len(a.arr) && a.arr[i] == value {
		return i
	}
	return -1
}

// Insert inserts value at correct index to preserve order.
//
// Time O(n) and space O(1).
func (a *OrderedArray[T]) Insert(value T) {
	i := bisect.BisectRight(a.arr, value)
	// Create space at index i and insert the value
	a.arr = append(a.arr[:i], append([]T{value}, a.arr[i:]...)...)
}

// Delete removes one occurrence of value and return it's index
// or returns -1.
//
// Time O(n) and space O(1).
func (a *OrderedArray[T]) Delete(value T) int {
	// Find the insertion point that would be to the right of any existing 'value's.
	// So, index `j-1` is the potential location of the rightmost 'value'.
	j := bisect.BisectRight(a.arr, value)
	i := j - 1

	// Check if 'i' is a valid index and if the element at 'i' is indeed the 'value'.
	// - i < 0: 'value' is smaller than all elements, or array is empty.
	// - a.arr[i] != value: The element at the found position is not 'value'.
	if i < 0 || i >= len(a.arr) || a.arr[i] != value {
		return -1 // Value not found
	}

	// Value found at index i (which is the rightmost occurrence)
	a.arr = append(a.arr[:i], a.arr[i+1:]...)
	return i
}

// Len returns OrderedArray's length.
func (a *OrderedArray[T]) Len() int {
	return len(a.arr)
}

// All returns an iterator over OrderedArray index-value pairs.
func (a *OrderedArray[T]) All() iter.Seq2[int, T] {
	return func(yield func(int, T) bool) {
		for i, v := range a.arr {
			if !yield(i, v) {
				return
			}
		}
	}
}
