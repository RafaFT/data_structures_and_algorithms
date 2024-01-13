package sets

import "fmt"

var _ Set[int] = &ArraySet[int]{}

// ArraySet is a naive [Set] implementation that uses an array underneath.
type ArraySet[T comparable] struct {
	arr []T
}

func (s *ArraySet[T]) String() string {
	return fmt.Sprintf("%v", s.arr)
}

func (s *ArraySet[T]) indexOf(value T) int {
	for i, v := range s.arr {
		if v == value {
			return i
		}
	}

	return -1
}

// Has reports whether value happens in ArraySet.
//
// Time O(n) and space O(1).
func (s *ArraySet[T]) Has(value T) bool {
	return s.indexOf(value) != -1
}

// Add adds value to ArraySet and reports whether it succeed.
//
// Time O(n) and space O(1).
func (s *ArraySet[T]) Add(value T) bool {
	if s.Has(value) {
		return false
	}

	s.arr = append(s.arr, value)

	return true
}

// Remove removes value from ArraySet and reports whether it was found.
//
// Time O(n) and space O(1).
func (s *ArraySet[T]) Remove(value T) bool {
	index := s.indexOf(value)

	if index == -1 {
		return false
	}

	copy(s.arr[index:len(s.arr)-1], s.arr[index+1:len(s.arr)])

	s.arr = s.arr[:len(s.arr)-1]

	return true
}
