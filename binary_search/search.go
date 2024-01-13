// Package binary_search exists to implement the common binary search algorithm
// and other similar related functions.
package binary_search

import (
	"cmp"
)

// Search returns the first index of s that contains value found by
// traditional binary search algorithm. Returns -1 if value is not found.
//
// s MUST be sorted in ascending order.
//
// Time O(log(n)) and space O(1).
func Search[T cmp.Ordered](s []T, value T) int {
	l, r := 0, len(s)-1

	for l <= r {
		m := (l + r) / 2

		if v := s[m]; v == value {
			return m
		} else if v < value {
			l = m + 1
		} else {
			r = m - 1
		}
	}

	return -1
}

// RightSearch returns the last index of s that either contains value
// or the index of s that value should be placed to keep order. It
// reports whether value was found or not.
// Returned index can be equal to len(s).
//
// s MUST be sorted in ascending order.
//
// Time O(log(n)) and space O(1).
func RightSearch[T cmp.Ordered](s []T, value T) (int, bool) {
	l, r, found := 0, len(s)-1, false

	for l <= r {
		m := (l + r) / 2

		if v := s[m]; v <= value {
			l = m + 1
			if v == value {
				found = true
			}
		} else {
			r = m - 1
		}
	}

	return l, found
}

// LeftSearch returns the first index of s that either contains value
// or the index of s that value should be placed to keep order. It
// reports whether value was found or not.
//
// s MUST be sorted in ascending order.
//
// Time O(log(n)) and space O(1).
func LeftSearch[T cmp.Ordered](s []T, value T) (int, bool) {
	l, r, found := 0, len(s)-1, false

	for l <= r {
		m := (l + r) / 2

		if v := s[m]; v >= value {
			r = m - 1
			if v == value {
				found = true
			}
		} else {
			l = m + 1
		}
	}

	return l, found
}
