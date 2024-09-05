package sets

import (
	"fmt"
	"reflect"
	"slices"
	"testing"
)

func TestArraySetHas(t *testing.T) {
	tests := []struct {
		set   *ArraySet[string]
		value string
		want  bool
	}{
		{
			&ArraySet[string]{arr: []string{}},
			"a",
			false,
		},
		{
			&ArraySet[string]{arr: []string{"a"}},
			"a",
			true,
		},
		{
			&ArraySet[string]{arr: []string{"a", "b", "c"}},
			"c",
			true,
		},
	}

	for i, test := range tests {
		if got := test.set.Has(test.value); got != test.want {
			t.Errorf("%d: %v.Has(%q) = %v, want %v", i, test.set, test.value, got, test.want)
		}
	}
}

func TestArraySetAdd(t *testing.T) {
	tests := []struct {
		set      *ArraySet[string]
		value    string
		wantBool bool
		wantSet  *ArraySet[string]
	}{
		{
			&ArraySet[string]{arr: []string{}},
			"a",
			true,
			&ArraySet[string]{arr: []string{"a"}},
		},
		{
			&ArraySet[string]{arr: []string{"a"}},
			"b",
			true,
			&ArraySet[string]{arr: []string{"a", "b"}},
		},
		{
			&ArraySet[string]{arr: []string{"a", "b"}},
			"a",
			false,
			&ArraySet[string]{arr: []string{"a", "b"}},
		},
	}

	for i, test := range tests {
		setBefore := fmt.Sprintf("%v", test.set.arr)

		if gotBool := test.set.Add(test.value); gotBool != test.wantBool || !reflect.DeepEqual(test.set, test.wantSet) {
			t.Errorf("%d: %v.Add(%q) = (%v, %v), want %v, %v",
				i, setBefore, test.value, test.set.arr, gotBool, test.wantSet.arr, test.wantBool)
		}
	}
}

func TestArraySetRemove(t *testing.T) {
	tests := []struct {
		set      *ArraySet[string]
		value    string
		wantBool bool
		wantSet  *ArraySet[string]
	}{
		{
			&ArraySet[string]{arr: []string{}},
			"a",
			false,
			&ArraySet[string]{arr: []string{}},
		},
		{
			&ArraySet[string]{arr: []string{"a"}},
			"b",
			false,
			&ArraySet[string]{arr: []string{"a"}},
		},
		{
			&ArraySet[string]{arr: []string{"a", "b"}},
			"a",
			true,
			&ArraySet[string]{arr: []string{"b"}},
		},
	}

	for i, test := range tests {
		setBefore := fmt.Sprintf("%v", test.set.arr)

		if gotBool := test.set.Remove(test.value); gotBool != test.wantBool || !reflect.DeepEqual(test.set, test.wantSet) {
			t.Errorf("%d: %v.Remove(%q) = (%v, %v), want %v, %v",
				i, setBefore, test.value, test.set.arr, gotBool, test.wantSet.arr, test.wantBool)
		}
	}
}

func TestArraySetValues(t *testing.T) {
	tests := []struct {
		array ArraySet[string]
		want  []string
	}{
		{
			ArraySet[string]{},
			nil,
		},
		{
			ArraySet[string]{
				arr: []string{},
			},
			nil,
		},
		{
			ArraySet[string]{
				arr: []string{"c", "b", "a"},
			},
			[]string{
				"c",
				"b",
				"a",
			},
		},
	}

	for i, test := range tests {
		if got := slices.Collect(test.array.Values()); !reflect.DeepEqual(got, test.want) {
			t.Errorf("%d: %s.Values() = %v, want %v", i, test.array, got, test.want)
		}
	}
}
