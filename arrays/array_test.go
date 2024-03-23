package arrays

import (
	"fmt"
	"reflect"
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
			Array[int]{1, 2, 3, 4},
			0,
			1,
		},
		{
			Array[int]{1, 2, 3, 4},
			3,
			4,
		},
	}

	for i, test := range tests {
		if got := test.array.Read(test.index); got != test.want {
			t.Errorf("%d: %v.Read(%d) = %d, want %d", i, test.array, test.index, got, test.want)
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
			Array[int]{1},
			-1,
		},
		{
			"index higher than length",
			Array[int]{1, 2, 3, 4},
			4,
		},
	}

	for i, test := range tests {
		if !panics(func() { test.array.Read(test.index) }) {
			t.Errorf("%d: %s should panic: %v.Read(%d)", i, test.name, test.array, test.index)
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
			Array[string]{},
			"a",
			-1,
		},
		{
			Array[string]{"c"},
			"b",
			-1,
		},
		{
			Array[string]{"a"},
			"a",
			0,
		},
		{
			Array[string]{"a", "b", "c", "c"},
			"c",
			2,
		},
		{
			Array[string]{"a", "b", "c", "d"},
			"d",
			3,
		},
	}

	for i, test := range tests {
		if got := test.array.Search(test.value); got != test.want {
			t.Errorf("%d: %v.Search(%q) = %d, want %d", i, test.array, test.value, got, test.want)
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
			Array[string]{},
			"a",
			0,
			Array[string]{"a"},
		},
		{
			Array[string]{"a"},
			"b",
			1,
			Array[string]{"a", "b"},
		},
		{
			Array[string]{"b"},
			"a",
			0,
			Array[string]{"a", "b"},
		},
		{
			Array[string]{"a", "c"},
			"b",
			1,
			Array[string]{"a", "b", "c"},
		},
	}

	for i, test := range tests {
		arrayBefore := fmt.Sprint(test.array)

		test.array.Insert(test.value, test.index)

		if !reflect.DeepEqual(test.array, test.want) {
			t.Errorf("%d: %s.Insert(%q, %d) = %s, want %s", i, arrayBefore, test.value, test.index, test.array, test.want)
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
			"index > 0 on empty array",
			Array[string]{},
			"a",
			1,
		},
		{
			"negative index",
			Array[string]{"a"},
			"b",
			-1,
		},
		{
			"index higher than length",
			Array[string]{"a"},
			"b",
			2,
		},
	}

	for i, test := range tests {
		arrayBefore := fmt.Sprint(test.array)

		if !panics(func() { test.array.Insert(test.value, test.index) }) {
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
			Array[int]{},
			1,
			Array[int]{},
			-1,
		},
		{
			Array[int]{1},
			2,
			Array[int]{1},
			-1,
		},
		{
			Array[int]{1},
			1,
			Array[int]{},
			0,
		},
		{
			Array[int]{1, 2},
			1,
			Array[int]{2},
			0,
		},
		{
			Array[int]{1, 2},
			2,
			Array[int]{1},
			1,
		},
		{
			Array[int]{0, 1, 2, 1},
			1,
			Array[int]{0, 2, 1},
			1,
		},
	}

	for i, test := range tests {
		arrayBefore := fmt.Sprint(test.array)

		gotIndex := test.array.Delete(test.value)

		if gotIndex != test.want || !reflect.DeepEqual(test.array, test.wantArray) {
			t.Errorf("%d: %s.Delete(%d) = (%v, %d), want (%v, %d)",
				i, arrayBefore, test.value, test.array, gotIndex, test.wantArray, test.want,
			)
		}
	}
}

func BenchmarkSearch(b *testing.B) {
	var a Array[int] = make([]int, 100000)

	for range b.N {
		a.Search(1)
	}
}

func BenchmarkInsert(b *testing.B) {
	var a Array[int] = make([]int, 1, 100000)

	for range b.N {
		a.Insert(0, len(a)-1)
	}
}

func BenchmarkDelete(b *testing.B) {
	var a Array[int] = make([]int, 100000)

	for range b.N {
		a.Delete(0)
	}
}
