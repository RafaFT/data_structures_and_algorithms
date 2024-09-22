package stacks

import "fmt"

var _ Stack[int] = &StackArray[int]{}

type StackArray[T any] struct {
	arr []T
}

func (s *StackArray[T]) Push(value T) {
	s.arr = append(s.arr, value)
}

func (s *StackArray[T]) Pop() (T, bool) {
	length := len(s.arr)
	if length == 0 {
		var v T
		return v, false
	}

	value := s.arr[length-1]
	s.arr = s.arr[:length-1]

	return value, true
}

func (s *StackArray[T]) Peek() (T, bool) {
	length := len(s.arr)
	if length == 0 {
		var v T
		return v, false
	}

	return s.arr[length-1], true
}

func (s *StackArray[T]) Len() int {
	return len(s.arr)
}

func (s *StackArray[T]) String() string {
	return fmt.Sprintf("StackArray{%v}", s.arr)
}
