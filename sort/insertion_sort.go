package sort

import (
	"cmp"

	"dsa/bisect"
)

// InsertionSort sorts s in-place by traversing the array
// once and finding the correct insertion position backwards
// in the array for each value.
// Because values at their correct position do not require
// insertion and swaps, this algorithm is very efficient for
// full and partially sorted arrays.
//
// Time O(N²) and space O(1).
func InsertionSort[T cmp.Ordered](s []T) {
	for i := 1; i < len(s); i++ {
		value := s[i]
		j := i - 1

		for j >= 0 && s[j] > value {
			s[j+1] = s[j]
			j--
		}

		s[j+1] = value
	}
}

// Like [InsertionSort], but uses bisect algorithm for
// finding insert position.
// It might be more efficient for worst-cases.
//
// Time O(N²) and space O(1).
func InsertionSortV2[T cmp.Ordered](s []T) {
	for i := 1; i < len(s); i++ {
		value := s[i]

		if j := bisect.BisectRight(s[:i], value); j != i {
			copy(s[j+1:i+1], s[j:i])
			s[j] = value
		}
	}
}
