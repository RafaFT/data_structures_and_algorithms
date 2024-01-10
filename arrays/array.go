package arrays

import (
	"fmt"
)

type Array[T comparable] struct {
	arr []T
}

func (a Array[T]) String() string {
	return fmt.Sprintf("%v", a.arr)
}

func Read[T comparable](a Array[T], index int) T {
	if index < 0 || index >= len(a.arr) {
		panic(fmt.Sprintf("index out of range [%d] with length %d", index, len(a.arr)))
	}

	return a.arr[index]
}

func Search[T comparable](a Array[T], value T) int {
	for i, v := range a.arr {
		if v == value {
			return i
		}
	}

	return -1
}

func Insert[T comparable](a Array[T], value T, index int) Array[T] {
	if index < 0 || index > len(a.arr) {
		panic(fmt.Sprintf("index out of range [%d] with length %d", index, len(a.arr)))
	}

	a.arr = append(a.arr, value)

	// swap values until new value is at correct index position
	for newIndex := len(a.arr) - 1; newIndex != index; newIndex-- {
		a.arr[newIndex], a.arr[newIndex-1] = a.arr[newIndex-1], a.arr[newIndex]
	}

	return a
}

func Delete[T comparable](a Array[T], value T) (Array[T], int) {
	index := Search(a, value)

	if index == -1 {
		return a, -1
	}

	// swap values to the right until last index can be removed
	for newIndex := index; newIndex < len(a.arr)-1; newIndex++ {
		a.arr[newIndex], a.arr[newIndex+1] = a.arr[newIndex+1], a.arr[newIndex]
	}

	a.arr = a.arr[:len(a.arr)-1]

	return a, index
}
