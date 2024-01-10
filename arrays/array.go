package arrays

import (
	"fmt"
)

type Array[T comparable] []T

func (a *Array[T]) Read(index int) T {
	arr := *a
	if index < 0 || index >= len(arr) {
		panic(fmt.Sprintf("index out of range [%d] with length %d", index, len(arr)))
	}

	return arr[index]
}

func (a *Array[T]) Search(value T) int {
	for i, v := range *a {
		if v == value {
			return i
		}
	}

	return -1
}

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
