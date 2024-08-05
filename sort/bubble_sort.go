package sort

import "cmp"

// BubbleSort implements the naive bubble sort algorithm.
// It traverses the array N times, swapping values when necessary,
// effectively moving the biggest value of each iteration to the end
// of the array, basically sorting it from right to left.
//
// Time O(N²) and space O(1).
func BubbleSort[T cmp.Ordered](s []T) {
	for i := 0; i < len(s)-1; i++ {
		// if no value was swapped the array is sorted
		swapped := false

		for j := 0; j < len(s)-1-i; j++ {
			if s[j] > s[j+1] {
				s[j], s[j+1] = s[j+1], s[j]
				swapped = true
			}
		}

		if !swapped {
			return
		}
	}
}
