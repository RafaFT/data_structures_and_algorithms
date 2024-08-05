package sort

import "cmp"

// SelectionSort implements the naive selection sort algorithm.
// It traverses the array N times, making at most one swap per iteration,
// moving the lowest value to the start of the array, basically sorting
// it from left to right.
//
// Time O(N²) and space O(1).
func SelectionSort[T cmp.Ordered](s []T) {
	for i, value := range s {
		lowest := value
		lowestIndex := i

		for j := i + 1; j < len(s); j++ {
			if nextValue := s[j]; nextValue < lowest {
				lowest = nextValue
				lowestIndex = j
			}
		}

		s[i], s[lowestIndex] = s[lowestIndex], s[i]
	}
}
