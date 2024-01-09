package arrays

import (
	"fmt"
	"math"
	"reflect"
	"strings"
	"testing"
)

func panics(f func()) (panicked bool) {
	defer func() {
		if e := recover(); e != nil {
			panicked = true
		}
	}()

	f()

	return panicked
}

func TestRead(t *testing.T) {
	tests := []struct {
		array Array[int]
		index int
		want  int
	}{
		{
			Array[int]{
				arr: []int{1, 2, 3, 4},
			},
			0,
			1,
		},
		{
			Array[int]{
				arr: []int{1, 2, 3, 4},
			},
			3,
			4,
		},
	}

	for i, test := range tests {
		if got := Read(test.array, test.index); got != test.want {
			t.Errorf("%d: Read(%v, %d) = %d, want %d", i, test.array.arr, test.index, got, test.want)
		}
	}
}

func TestReadError(t *testing.T) {
	tests := []struct {
		name  string
		array Array[int]
		index int
	}{
		{
			"read from empty array",
			Array[int]{},
			0,
		},
		{
			"negative index",
			Array[int]{arr: []int{1}},
			-1,
		},
		{
			"index higher than length",
			Array[int]{
				arr: []int{1, 2, 3, 4},
			},
			4,
		},
	}

	for i, test := range tests {
		if !panics(func() { Read(test.array, test.index) }) {
			t.Errorf("%d: %s should panic: Read(%v, %d)", i, test.name, test.array.arr, test.index)
		}
	}
}

func TestSearch(t *testing.T) {
	tests := []struct {
		array Array[string]
		value string
		want  int
	}{
		{
			Array[string]{arr: []string{}},
			"a",
			-1,
		},
		{
			Array[string]{arr: []string{"c"}},
			"b",
			-1,
		},
		{
			Array[string]{arr: []string{"a"}},
			"a",
			0,
		},
		{
			Array[string]{arr: []string{"a", "b", "c", "c"}},
			"c",
			2,
		},
		{
			Array[string]{arr: []string{"a", "b", "c", "d"}},
			"d",
			3,
		},
	}

	for i, test := range tests {
		if got := Search(test.array, test.value); got != test.want {
			t.Errorf("%d: Search(%v, %q) = %d, want %d", i, test.array.arr, test.value, got, test.want)
		}
	}
}

func TestSearchFunc(t *testing.T) {
	tests := []struct {
		array Array[string]
		value string
		want  int
	}{
		{
			Array[string]{arr: []string{}},
			"A",
			-1,
		},
		{
			Array[string]{arr: []string{"c"}},
			"B",
			-1,
		},
		{
			Array[string]{arr: []string{"a"}},
			"A",
			0,
		},
		{
			Array[string]{arr: []string{"a", "b", "c", "c"}},
			"C",
			2,
		},
		{
			Array[string]{arr: []string{"a", "b", "c", "d"}},
			"D",
			3,
		},
	}

	cmp := func(s1, s2 string) bool {
		return strings.EqualFold(s1, s2)
	}

	for i, test := range tests {

		if got := SearchFunc(test.array, test.value, cmp); got != test.want {
			t.Errorf("%d: Search(%v, %q) = %d, want %d", i, test.array.arr, test.value, got, test.want)
		}
	}
}

func TestInsert(t *testing.T) {
	tests := []struct {
		array Array[string]
		value string
		index int
		want  Array[string]
	}{
		{
			Array[string]{arr: []string{}},
			"a",
			0,
			Array[string]{arr: []string{"a"}},
		},
		{
			Array[string]{arr: []string{"a"}},
			"b",
			1,
			Array[string]{arr: []string{"a", "b"}},
		},
		{
			Array[string]{arr: []string{"b"}},
			"a",
			0,
			Array[string]{arr: []string{"a", "b"}},
		},
		{
			Array[string]{arr: []string{"a", "c"}},
			"b",
			1,
			Array[string]{arr: []string{"a", "b", "c"}},
		},
	}

	for i, test := range tests {
		arrayBefore := fmt.Sprintf("%v", test.array.arr)

		if got := Insert(test.array, test.value, test.index); !reflect.DeepEqual(got, test.want) {
			t.Errorf("%d: Insert(%s, %q, %d) = %v, want %v", i, arrayBefore, test.value, test.index, test.array.arr, test.want.arr)
		}
	}
}

func TestInsertError(t *testing.T) {
	tests := []struct {
		name  string
		array Array[string]
		value string
		index int
	}{
		{
			"negative index",
			Array[string]{arr: []string{"a"}},
			"b",
			-1,
		},
		{
			"index higher than length",
			Array[string]{arr: []string{"a"}},
			"b",
			2,
		},
	}

	for i, test := range tests {
		arrayBefore := fmt.Sprintf("%v", test.array.arr)

		if !panics(func() { Insert(test.array, test.value, test.index) }) {
			t.Errorf("%d: %s should panic: %s.Insert(%q, %d)", i, test.name, arrayBefore, test.value, test.index)
		}
	}
}

func TestDelete(t *testing.T) {
	tests := []struct {
		array     Array[int]
		value     int
		wantArray Array[int]
		want      int
	}{
		{
			Array[int]{arr: []int{}},
			1,
			Array[int]{arr: []int{}},
			-1,
		},
		{
			Array[int]{arr: []int{1}},
			2,
			Array[int]{arr: []int{1}},
			-1,
		},
		{
			Array[int]{arr: []int{1}},
			1,
			Array[int]{arr: []int{}},
			0,
		},
		{
			Array[int]{arr: []int{1, 2}},
			1,
			Array[int]{arr: []int{2}},
			0,
		},
		{
			Array[int]{arr: []int{1, 2}},
			2,
			Array[int]{arr: []int{1}},
			1,
		},
		{
			Array[int]{arr: []int{0, 1, 2, 1}},
			1,
			Array[int]{arr: []int{0, 2, 1}},
			1,
		},
	}

	for i, test := range tests {
		arrayBefore := fmt.Sprintf("%v", test.array.arr)

		gotArray, gotIndex := Delete(test.array, test.value)

		if gotIndex != test.want || !reflect.DeepEqual(gotArray.arr, test.wantArray.arr) {
			t.Errorf("%d: Delete(%s, %d) = (%v, %d), want (%v, %d)",
				i, arrayBefore, test.value, test.array.arr, gotIndex, test.wantArray.arr, test.want,
			)
		}
	}
}

func TestDeleteFunc(t *testing.T) {
	tests := []struct {
		array     Array[int]
		value     int
		wantArray Array[int]
		want      int
	}{
		{
			Array[int]{arr: []int{}},
			-1,
			Array[int]{arr: []int{}},
			-1,
		},
		{
			Array[int]{arr: []int{1}},
			-2,
			Array[int]{arr: []int{1}},
			-1,
		},
		{
			Array[int]{arr: []int{1}},
			-1,
			Array[int]{arr: []int{}},
			0,
		},
		{
			Array[int]{arr: []int{1, 2}},
			-1,
			Array[int]{arr: []int{2}},
			0,
		},
		{
			Array[int]{arr: []int{1, 2}},
			-2,
			Array[int]{arr: []int{1}},
			1,
		},
		{
			Array[int]{arr: []int{0, 1, 2, 1}},
			-1,
			Array[int]{arr: []int{0, 2, 1}},
			1,
		},
	}

	cmp := func(v1, v2 int) bool {
		return math.Abs(float64(v1)) == math.Abs(float64(v2))
	}

	for i, test := range tests {
		arrayBefore := fmt.Sprintf("%v", test.array.arr)

		gotArray, gotIndex := DeleteFunc(test.array, test.value, cmp)

		if gotIndex != test.want || !reflect.DeepEqual(gotArray.arr, test.wantArray.arr) {
			t.Errorf("%d: Delete(%s, %d) = (%v, %d), want (%v, %d)",
				i, arrayBefore, test.value, test.array.arr, gotIndex, test.wantArray.arr, test.want,
			)
		}
	}
}
