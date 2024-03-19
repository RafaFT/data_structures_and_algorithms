// Package bisect exists to implement the common binary search algorithm
// and other similar related functions.
package bisect

import (
	"cmp"
)

// Search returns index of s that contains v. It searches using
// traditional binary search implementation and returned index
// is not guarantee to be either first or last occurrence of v.
// It returns -1 if v is not found.
//
// s MUST be sorted in ascending order.
//
// Time O(log(n)) and space O(1).
func Search[T cmp.Ordered](s []T, v T) int {
	left, right := 0, len(s)-1

	for left <= right {
		m := ((right - left) / 2) + left
		value := s[m]

		if value == v {
			return m
		}

		if value > v {
			right = m - 1
		} else {
			left = m + 1
		}
	}

	return -1
}

// BisectLeft returns the lowest index of s where v should be placed to keep order.
// Index might equal len(s).
//
// s MUST be sorted in ascending order.
//
// Time O(log(n)) and space O(1).
func BisectLeft[T cmp.Ordered](s []T, value T) int {
	left, right := 0, len(s)-1

	for left <= right {
		m := ((right - left) / 2) + left
		if v := s[m]; v >= value {
			right = m - 1
		} else {
			left = m + 1
		}
	}

	return left
}

// BisectRight returns the highest index of s where v should be placed to keep order.
// Index might equal len(s).
//
// s MUST be sorted in ascending order.
//
// Time O(log(n)) and space O(1).
func BisectRight[T cmp.Ordered](s []T, v T) int {
	left, right := 0, len(s)-1

	for left <= right {
		m := ((right - left) / 2) + left
		if value := s[m]; v >= value {
			left = m + 1
		} else {
			right = m - 1
		}
	}

	return left
}
