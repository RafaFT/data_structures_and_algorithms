package sets

import (
	"fmt"
	"reflect"
	"testing"
)

func TestArraySetHas(t *testing.T) {
	tests := []struct {
		set   *arraySet[string]
		value string
		want  bool
	}{
		{
			&arraySet[string]{arr: []string{}},
			"a",
			false,
		},
		{
			&arraySet[string]{arr: []string{"a"}},
			"a",
			true,
		},
		{
			&arraySet[string]{arr: []string{"a", "b", "c"}},
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
		set      *arraySet[string]
		value    string
		wantBool bool
		wantSet  *arraySet[string]
	}{
		{
			&arraySet[string]{arr: []string{}},
			"a",
			true,
			&arraySet[string]{arr: []string{"a"}},
		},
		{
			&arraySet[string]{arr: []string{"a"}},
			"b",
			true,
			&arraySet[string]{arr: []string{"a", "b"}},
		},
		{
			&arraySet[string]{arr: []string{"a", "b"}},
			"a",
			false,
			&arraySet[string]{arr: []string{"a", "b"}},
		},
	}

	for i, test := range tests {
		setBefore := fmt.Sprint(test.set)

		if gotBool := test.set.Add(test.value); gotBool != test.wantBool || !reflect.DeepEqual(test.set, test.wantSet) {
			t.Errorf("%d: %v.Add(%q) = (%v, %v), want %v, %v",
				i, setBefore, test.value, test.set, gotBool, test.wantSet, test.wantBool)
		}
	}
}

func TestArraySetRemove(t *testing.T) {
	tests := []struct {
		set      *arraySet[string]
		value    string
		wantBool bool
		wantSet  *arraySet[string]
	}{
		{
			&arraySet[string]{arr: []string{}},
			"a",
			false,
			&arraySet[string]{arr: []string{}},
		},
		{
			&arraySet[string]{arr: []string{"a"}},
			"b",
			false,
			&arraySet[string]{arr: []string{"a",}},
		},
		{
			&arraySet[string]{arr: []string{"a", "b"}},
			"a",
			true,
			&arraySet[string]{arr: []string{"b"}},
		},
	}

	for i, test := range tests {
		setBefore := fmt.Sprint(test.set)

		if gotBool := test.set.Remove(test.value); gotBool != test.wantBool || !reflect.DeepEqual(test.set, test.wantSet) {
			t.Errorf("%d: %v.Remove(%q) = (%v, %v), want %v, %v",
				i, setBefore, test.value, test.set, gotBool, test.wantSet, test.wantBool)
		}
	}
}
