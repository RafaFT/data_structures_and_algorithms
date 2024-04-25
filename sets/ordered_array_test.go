package sets

import (
	"fmt"
	"reflect"
	"testing"
)

func TestOrderedArraySetHas(t *testing.T) {
	tests := []struct {
		set   *OrderedArraySet[string]
		value string
		want  bool
	}{
		{
			&OrderedArraySet[string]{arr: []string{}},
			"a",
			false,
		},
		{
			&OrderedArraySet[string]{arr: []string{"a"}},
			"a",
			true,
		},
		{
			&OrderedArraySet[string]{arr: []string{"a", "b", "c"}},
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

func TestOrderedArraySetAdd(t *testing.T) {
	tests := []struct {
		set      *OrderedArraySet[string]
		value    string
		wantBool bool
		wantSet  *OrderedArraySet[string]
	}{
		{
			&OrderedArraySet[string]{arr: []string{}},
			"a",
			true,
			&OrderedArraySet[string]{arr: []string{"a"}},
		},
		{
			&OrderedArraySet[string]{arr: []string{"a"}},
			"b",
			true,
			&OrderedArraySet[string]{arr: []string{"a", "b"}},
		},
		{
			&OrderedArraySet[string]{arr: []string{"a", "b"}},
			"a",
			false,
			&OrderedArraySet[string]{arr: []string{"a", "b"}},
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

func TestOrderedArraySetRemove(t *testing.T) {
	tests := []struct {
		set      *OrderedArraySet[string]
		value    string
		wantBool bool
		wantSet  *OrderedArraySet[string]
	}{
		{
			&OrderedArraySet[string]{arr: []string{}},
			"a",
			false,
			&OrderedArraySet[string]{arr: []string{}},
		},
		{
			&OrderedArraySet[string]{arr: []string{"a"}},
			"b",
			false,
			&OrderedArraySet[string]{arr: []string{"a"}},
		},
		{
			&OrderedArraySet[string]{arr: []string{"a", "b"}},
			"a",
			true,
			&OrderedArraySet[string]{arr: []string{"b"}},
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
