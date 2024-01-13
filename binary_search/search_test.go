package binary_search

import "testing"

func TestSearch(t *testing.T) {
	tests := []struct {
		s    []string
		v    string
		want int
	}{
		{
			[]string{},
			"c",
			-1,
		},
		{
			[]string{"z"},
			"a",
			-1,
		},
		{
			[]string{"a", "b", "d"},
			"a",
			0,
		},
		{
			[]string{"a", "b", "d"},
			"b",
			1,
		},
		{
			[]string{"a", "b", "d"},
			"d",
			2,
		},
		{
			[]string{"a", "b", "d"},
			"c",
			-1,
		},
	}

	for i, test := range tests {
		if got := Search(test.s, test.v); got != test.want {
			t.Errorf("%d: Search(%v, %v) = %d, want %d", i, test.s, test.v, got, test.want)
		}
	}
}

func TestSearchRight(t *testing.T) {
	tests := []struct {
		s        []string
		v        string
		wantInt  int
		wantBool bool
	}{
		{
			[]string{},
			"a",
			0,
			false,
		},
		{
			[]string{"b"},
			"a",
			0,
			false,
		},
		{
			[]string{"b"},
			"c",
			1,
			false,
		},
		{
			[]string{"b"},
			"b",
			1,
			true,
		},
		{
			[]string{"a", "c", "c"},
			"a",
			1,
			true,
		},
		{
			[]string{"a", "c", "c"},
			"b",
			1,
			false,
		},
		{
			[]string{"a", "c", "c"},
			"c",
			3,
			true,
		},
		{
			[]string{"b", "b", "b"},
			"b",
			3,
			true,
		},
		{
			[]string{"b", "b", "b"},
			"a",
			0,
			false,
		},
	}

	for i, test := range tests {
		if gotInt, gotBool := RightSearch(test.s, test.v); gotInt != test.wantInt || gotBool != test.wantBool {
			t.Errorf("%d: SearchRight(%v, %v) = (%d, %v), want (%d, %v)", i, test.s, test.v, gotInt, gotBool, test.wantInt, test.wantBool)
		}
	}
}

func TestSearchLeft(t *testing.T) {
	tests := []struct {
		s        []string
		v        string
		wantInt  int
		wantBool bool
	}{
		{
			[]string{},
			"a",
			0,
			false,
		},
		{
			[]string{"b"},
			"a",
			0,
			false,
		},
		{
			[]string{"b"},
			"c",
			1,
			false,
		},
		{
			[]string{"b"},
			"b",
			0,
			true,
		},
		{
			[]string{"a", "c", "c"},
			"a",
			0,
			true,
		},
		{
			[]string{"a", "c", "c"},
			"b",
			1,
			false,
		},
		{
			[]string{"a", "c", "c"},
			"c",
			1,
			true,
		},
		{
			[]string{"b", "b", "b"},
			"b",
			0,
			true,
		},
		{
			[]string{"b", "b", "b"},
			"a",
			0,
			false,
		},
	}

	for i, test := range tests {
		if gotInt, gotBool := LeftSearch(test.s, test.v); gotInt != test.wantInt || gotBool != test.wantBool {
			t.Errorf("%d: SearchLeft(%v, %v) = (%d, %v), want (%d, %v)", i, test.s, test.v, gotInt, gotBool, test.wantInt, test.wantBool)
		}
	}
}
