package arrays

import (
	"fmt"
	"iter"
	"reflect"
	"testing"
)

type collected2[T1, T2 any] struct {
	v1 T1
	v2 T2
}

func collect[T any](seq iter.Seq[T]) []T {
	var s []T

	for v := range seq {
		s = append(s, v)
	}

	return s
}

func collect2[T1, T2 any](seq iter.Seq2[T1, T2]) []collected2[T1, T2] {
	var s []collected2[T1, T2]

	for v1, v2 := range seq {
		s = append(s, collected2[T1, T2]{v1, v2})
	}

	return s
}

func TestOrderedRead(t *testing.T) {
	tests := []struct {
		array OrderedArray[int]
		index int
		want  int
	}{
		{
			OrderedArray[int]{arr: []int{1, 2, 3, 4}},
			0,
			1,
		},
		{
			OrderedArray[int]{arr: []int{1, 2, 3, 4}},
			3,
			4,
		},
	}

	for i, test := range tests {
		if got := test.array.Read(test.index); got != test.want {
			t.Errorf("%d: %v.Read(%d) = %d, want %d", i, test.array.arr, test.index, got, test.want)
		}
	}
}

func TestOrderedReadError(t *testing.T) {
	tests := []struct {
		name  string
		array OrderedArray[int]
		index int
	}{
		{
			"read from empty array",
			OrderedArray[int]{},
			0,
		},
		{
			"negative index",
			OrderedArray[int]{arr: []int{1}},
			-1,
		},
		{
			"index higher than length",
			OrderedArray[int]{arr: []int{1, 2, 3, 4}},
			4,
		},
	}

	for i, test := range tests {
		if !panics(func() { test.array.Read(test.index) }) {
			t.Errorf("%d: %s should panic: %v.Read(%d)", i, test.name, test.array.arr, test.index)
		}
	}
}

func TestOrderedSearch(t *testing.T) {
	tests := []struct {
		array OrderedArray[string]
		value string
		want  int
	}{
		{
			OrderedArray[string]{},
			"a",
			-1,
		},
		{
			OrderedArray[string]{arr: []string{"c"}},
			"b",
			-1,
		},
		{
			OrderedArray[string]{arr: []string{"a"}},
			"a",
			0,
		},
		{
			OrderedArray[string]{arr: []string{"a", "b", "c", "c"}},
			"c",
			2,
		},
		{
			OrderedArray[string]{arr: []string{"a", "b", "c", "d"}},
			"d",
			3,
		},
		{
			OrderedArray[string]{arr: []string{"a", "b", "c", "d"}},
			"e",
			-1,
		},
		{
			OrderedArray[string]{arr: []string{"a", "a", "a", "a"}},
			"a",
			0,
		},
	}

	for i, test := range tests {
		if got := test.array.Search(test.value); got != test.want {
			t.Errorf("%d: %v.Search(%q) = %d, want %d", i, test.array.arr, test.value, got, test.want)
		}
	}
}

func TestOrderedInsert(t *testing.T) {
	tests := []struct {
		array OrderedArray[string]
		value string
		want  OrderedArray[string]
	}{
		{
			OrderedArray[string]{arr: []string{}},
			"a",
			OrderedArray[string]{arr: []string{"a"}},
		},
		{
			OrderedArray[string]{arr: []string{"a"}},
			"b",
			OrderedArray[string]{arr: []string{"a", "b"}},
		},
		{
			OrderedArray[string]{arr: []string{"b"}},
			"a",
			OrderedArray[string]{arr: []string{"a", "b"}},
		},
		{
			OrderedArray[string]{arr: []string{"a", "c"}},
			"b",
			OrderedArray[string]{arr: []string{"a", "b", "c"}},
		},
	}

	for i, test := range tests {
		arrayBefore := fmt.Sprint(test.array.arr)

		test.array.Insert(test.value)

		if !reflect.DeepEqual(test.array, test.want) {
			t.Errorf("%d: %s.Insert(%q) = %s, want %s", i, arrayBefore, test.value, test.array.arr, test.want.arr)
		}
	}
}

func TestOrderedDelete(t *testing.T) {
	tests := []struct {
		array     OrderedArray[int]
		value     int
		wantArray OrderedArray[int]
		want      int
	}{
		{
			OrderedArray[int]{arr: []int{}},
			1,
			OrderedArray[int]{arr: []int{}},
			-1,
		},
		{
			OrderedArray[int]{arr: []int{1}},
			2,
			OrderedArray[int]{arr: []int{1}},
			-1,
		},
		{
			OrderedArray[int]{arr: []int{1}},
			1,
			OrderedArray[int]{arr: []int{}},
			0,
		},
		{
			OrderedArray[int]{arr: []int{1, 1, 2, 2}},
			1,
			OrderedArray[int]{arr: []int{1, 2, 2}},
			1,
		},
		{
			OrderedArray[int]{arr: []int{1, 1, 2, 2}},
			2,
			OrderedArray[int]{arr: []int{1, 1, 2}},
			3,
		},
	}

	for i, test := range tests {
		arrayBefore := fmt.Sprint(test.array.arr)

		gotIndex := test.array.Delete(test.value)

		if gotIndex != test.want || !reflect.DeepEqual(test.array, test.wantArray) {
			t.Errorf("%d: %s.Delete(%d) = (%v, %d), want (%v, %d)",
				i, arrayBefore, test.value, test.array.arr, gotIndex, test.wantArray.arr, test.want,
			)
		}
	}
}

func BenchmarkOrderedSearch(b *testing.B) {
	oa := OrderedArray[int]{
		arr: make([]int, 100000),
	}

	for range b.N {
		oa.Search(1)
	}
}

func BenchmarkOrderedInsert(b *testing.B) {
	oa := OrderedArray[int]{
		arr: make([]int, 0, 100000),
	}

	for range b.N {
		oa.Insert(0)
	}
}

func BenchmarkOrderedDelete(b *testing.B) {
	oa := OrderedArray[int]{
		arr: make([]int, 100000),
	}

	for range b.N {
		oa.Delete(0)
	}
}

func TestOrderedArrayAll(t *testing.T) {
	oa := OrderedArray[int]{
		arr: []int{
			1,
			2,
			3,
		},
	}

	want := []collected2[int, int]{
		collected2[int, int]{0, 1},
		collected2[int, int]{1, 2},
		collected2[int, int]{2, 3},
	}

	if got := collect2(oa.All()); !reflect.DeepEqual(got, want) {
		t.Errorf("oa.All() = %v, want %v", got, want)
	}
}

func TestOrderedArrayValues(t *testing.T) {
	oa := OrderedArray[int]{
		arr: []int{
			1,
			2,
			3,
		},
	}

	want := []int{
		1,
		2,
		3,
	}

	if got := collect(oa.Values()); !reflect.DeepEqual(got, want) {
		t.Errorf("oa.Values() = %v, want %v", got, want)
	}
}

func TestOrderedArrayOccurrences(t *testing.T) {
	tests := []struct {
		oa   OrderedArray[int]
		want []collected2[int, int]
	}{
		{
			OrderedArray[int]{},
			nil,
		},
		{
			OrderedArray[int]{
				arr: []int{
					1,
				},
			},
			[]collected2[int, int]{
				collected2[int, int]{1, 1},
			},
		},
		{
			OrderedArray[int]{
				arr: []int{
					1,
					1,
					1,
				},
			},
			[]collected2[int, int]{
				collected2[int, int]{1, 3},
			},
		},
		{
			OrderedArray[int]{
				arr: []int{
					1,
					2,
					3,
				},
			},
			[]collected2[int, int]{
				collected2[int, int]{1, 1},
				collected2[int, int]{2, 1},
				collected2[int, int]{3, 1},
			},
		},
		{
			OrderedArray[int]{
				arr: []int{
					1,
					2,
					2,
					3,
				},
			},
			[]collected2[int, int]{
				collected2[int, int]{1, 1},
				collected2[int, int]{2, 2},
				collected2[int, int]{3, 1},
			},
		},
		{
			OrderedArray[int]{
				arr: []int{
					1,
					1,
					2,
					3,
				},
			},
			[]collected2[int, int]{
				collected2[int, int]{1, 2},
				collected2[int, int]{2, 1},
				collected2[int, int]{3, 1},
			},
		},
		{
			OrderedArray[int]{
				arr: []int{
					1,
					2,
					3,
					3,
				},
			},
			[]collected2[int, int]{
				collected2[int, int]{1, 1},
				collected2[int, int]{2, 1},
				collected2[int, int]{3, 2},
			},
		},
	}

	for i, test := range tests {
		if got := collect2(test.oa.Occurrences()); !reflect.DeepEqual(got, test.want) {
			t.Errorf("%d: %v.Occurrences() = %v, want %v", i, test.oa.arr, got, test.want)
		}
	}
}
