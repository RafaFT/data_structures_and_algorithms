// Package arrays contains all array's implementations, which are defined as
// concrete data structures that stores elements sequentially in memory.
//
// Since it's not possible to implement an array from scratch in Go, all current
// implementations use slices under the hood and define all necessary behavior through
// functions/methods.
//
// Most if not all functionalities defined here are also implemented by the new "slices"
// package, which makes this package and it's types mostly useless.
package arrays

import (
	"fmt"
)

// Array is a simple mutable array implementation.
type Array[T comparable] []T

// Read returns the value at the provided index.
// It panics if index < 0 or index >= len(a).
//
// Time O(1) and space O(1).
func (a *Array[T]) Read(index int) T {
	arr := *a
	if index < 0 || index >= len(arr) {
		panic(fmt.Sprintf("index out of range [%d] with length %d", index, len(arr)))
	}

	return arr[index]
}

// Search returns the first index that contains value or -1.
//
// Time O(n) and space O(1).
func (a *Array[T]) Search(value T) int {
	for i, v := range *a {
		if v == value {
			return i
		}
	}

	return -1
}

// Insert inserts value at the provided index.
// It panics if index < 0 or index > len(a)
//
// Time O(n) and space O(1).
func (a *Array[T]) Insert(value T, index int) {
	arr := *a
	if index < 0 || index > len(arr) {
		panic(fmt.Sprintf("index out of range [%d] with length %d", index, len(arr)))
	}

	arr = append(arr, value)

	// swap values until new value is at correct index position
	for newIndex := len(arr) - 1; newIndex != index; newIndex-- {
		arr[newIndex], arr[newIndex-1] = arr[newIndex-1], arr[newIndex]
	}

	*a = arr
}

// Delete removes the first occurrence of value and return it's index or returns -1.
//
// Time O(n) and space O(1).
func (a *Array[T]) Delete(value T) int {
	arr := *a

	index := arr.Search(value)

	if index == -1 {
		return -1
	}

	// swap values to the right until last index can be removed
	for newIndex := index; newIndex < len(arr)-1; newIndex++ {
		arr[newIndex], arr[newIndex+1] = arr[newIndex+1], arr[newIndex]
	}

	*a = arr[:len(arr)-1]

	return index
}
