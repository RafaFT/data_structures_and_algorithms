package sets

import (
	"cmp"
	"iter"

	"dsa/bisect"
)

var _ Set[int] = &OrderedArraySet[int]{}

// OrderedArraySet is a [Set] implementation that uses an ordered array underneath.
type OrderedArraySet[T cmp.Ordered] struct {
	arr []T
}

// Has reports whether value happens in OrderedArray.
//
// Time O(log(n)) and space O(1).
func (s *OrderedArraySet[T]) Has(value T) bool {
	return bisect.Search(s.arr, value) != -1
}

// Add adds value to OrderedArray and reports whether it succeed.
//
// Time O(n) and space O(1).
func (s *OrderedArraySet[T]) Add(value T) bool {
	i := bisect.BisectLeft(s.arr, value)

	if i == len(s.arr) {
		s.arr = append(s.arr, value)
		return true
	}

	if s.arr[i] == value {
		return false
	}

	// appended value doesn't matter
	s.arr = append(s.arr, value)
	copy(s.arr[i+1:], s.arr[i:len(s.arr)-1])
	s.arr[i] = value

	return true
}

// Remove removes value from OrderedArray and reports whether it was found.
//
// Time O(n) and space O(1).
func (s *OrderedArraySet[T]) Remove(value T) bool {
	i := bisect.BisectLeft(s.arr, value)

	if i == len(s.arr) || s.arr[i] != value {
		return false
	}

	copy(s.arr[i:len(s.arr)-1], s.arr[i+1:])
	s.arr = s.arr[:len(s.arr)-1]

	return true
}

// Values returns an iterator of OrderedArraySet elements.
func (s *OrderedArraySet[T]) Values() iter.Seq[T] {
	return func(yield func(T) bool) {
		for _, value := range s.arr {
			if !yield(value) {
				return
			}
		}
	}
}
