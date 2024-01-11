package sets

import "fmt"

var _ Set[int] = &arraySet[int]{}

type arraySet[T comparable] struct {
	arr []T
}

func (s *arraySet[T]) String() string {
	return fmt.Sprintf("%v", s.arr)
}

func (s *arraySet[T]) indexOf(value T) int {
	for i, v := range s.arr {
		if v == value {
			return i
		}
	}

	return -1
}

func (s *arraySet[T]) Has(value T) bool {
	return s.indexOf(value) != -1
}

func (s *arraySet[T]) Add(value T) bool {
	if s.Has(value) {
		return false
	}

	s.arr = append(s.arr, value)

	return true
}

func (s *arraySet[T]) Remove(value T) bool {
	index := s.indexOf(value)

	if index == -1 {
		return false
	}

	copy(s.arr[index:len(s.arr)-1], s.arr[index+1:len(s.arr)])

	s.arr = s.arr[:len(s.arr)-1]

	return true
}
