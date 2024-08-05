package sort

import (
	"fmt"
	"slices"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	tests := []struct {
		arr  []int
		want []int
	}{
		{
			nil,
			nil,
		},
		{
			[]int{},
			[]int{},
		},
		{
			[]int{2},
			[]int{2},
		},
		{
			[]int{1, 2, 3},
			[]int{1, 2, 3},
		},
		{
			[]int{3, 2, 1},
			[]int{1, 2, 3},
		},
		{
			[]int{0, -2, 2, -3, 3},
			[]int{-3, -2, 0, 2, 3},
		},
	}

	for i, test := range tests {
		arrStr := fmt.Sprint(test.arr)

		if BubbleSort(test.arr); slices.Compare(test.arr, test.want) != 0 {
			t.Errorf("%d: BubbleSort(%s) = %v, want %v", i, arrStr, test.arr, test.want)
		}
	}
}

func BenchmarkBubbleSort(b *testing.B) {
	const length = 1000
	testCopy := make([]int, length)

	b.Run(fmt.Sprintf("best case size = %d", length), func(b *testing.B) {
		original := make([]int, 0, length)

		for i := range length {
			original = append(original, i)
		}

		b.ResetTimer()
		b.ReportAllocs()
		for range b.N {
			copy(testCopy, original)
			BubbleSort(testCopy)
		}
	})

	b.Run(fmt.Sprintf("worst case size = %d", length), func(b *testing.B) {
		original := make([]int, 0, length)

		for i := length - 1; i >= 0; i-- {
			original = append(original, i)
		}

		b.ResetTimer()
		b.ReportAllocs()
		for range b.N {
			copy(testCopy, original)
			BubbleSort(testCopy)
		}
	})
}
