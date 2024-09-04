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

	if i == len(a.arr) || i == 0 && a.arr[i] != value {
		return -1
	}

	return i
}

// Insert inserts value at correct index to preserve order.
//
// Time O(n) and space O(1).
func (a *OrderedArray[T]) Insert(value T) {
	i := bisect.BisectRight(a.arr, value)

	a.arr = append(a.arr, value)

	if i != -1 {
		for ; i < len(a.arr)-1; i++ {
			a.arr[i], a.arr[i+1] = a.arr[i+1], a.arr[i]
		}
	} else {
		for i := len(a.arr) - 1; i > 0 && value < a.arr[i-1]; i-- {
			a.arr[i], a.arr[i-1] = a.arr[i-1], a.arr[i]
		}
	}
}

// Delete removes one occurrence of value and return it's index
// or returns -1.
//
// Time O(n) and space O(1).
func (a *OrderedArray[T]) Delete(value T) int {
	i := bisect.BisectRight(a.arr, value) - 1

	if i == -1 || (a.arr[i] != value) {
		return -1
	}

	for j := i; j < len(a.arr)-1; j++ {
		a.arr[j], a.arr[j+1] = a.arr[j+1], a.arr[j]
	}

	a.arr = a.arr[:len(a.arr)-1]

	return i
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
