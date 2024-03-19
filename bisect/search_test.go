package bisect

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
			[]string{"b", "d", "f"},
			"a",
			-1,
		},
		{
			[]string{"b", "d", "f"},
			"b",
			0,
		},
		{
			[]string{"b", "d", "f"},
			"c",
			-1,
		},
		{
			[]string{"b", "d", "f"},
			"d",
			1,
		},
		{
			[]string{"b", "d", "f"},
			"e",
			-1,
		},
		{
			[]string{"b", "d", "f"},
			"f",
			2,
		},
		{
			[]string{"b", "d", "f"},
			"g",
			-1,
		},
	}

	for i, test := range tests {
		if got := Search(test.s, test.v); got != test.want {
			t.Errorf("%d: Search(%v, %v) = %d, want %d", i, test.s, test.v, got, test.want)
		}
	}
}

func TestBisectLeft(t *testing.T) {
	tests := []struct {
		s    []string
		v    string
		want int
	}{
		{
			[]string{},
			"a",
			0,
		},
		{
			[]string{"b"},
			"a",
			0,
		},
		{
			[]string{"b"},
			"b",
			0,
		},
		{
			[]string{"b"},
			"c",
			1,
		},
		{
			[]string{"b", "d", "f"},
			"a",
			0,
		},
		{
			[]string{"b", "d", "f"},
			"b",
			0,
		},
		{
			[]string{"b", "d", "f"},
			"c",
			1,
		},
		{
			[]string{"b", "d", "f"},
			"d",
			1,
		},
		{
			[]string{"b", "d", "f"},
			"e",
			2,
		},
		{
			[]string{"b", "d", "f"},
			"f",
			2,
		},
		{
			[]string{"b", "d", "f"},
			"g",
			3,
		},
		{
			[]string{"b", "b", "b"},
			"b",
			0,
		},
	}

	for i, test := range tests {
		if gotInt := BisectLeft(test.s, test.v); gotInt != test.want {
			t.Errorf("%d: BisectLeft(%v, %v) = %d, want %d", i, test.s, test.v, gotInt, test.want)
		}
	}
}
