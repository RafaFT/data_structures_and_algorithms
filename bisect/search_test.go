package bisect

import "testing"

func TestSearch(t *testing.T) {
	tests := []struct {
		s    []string
		v    string
		want int // For Search, if element is found, 'want' is an example index. If not found, 'want' is -1.
	}{
		{[]string{}, "c", -1},
		{[]string{"b", "d", "f"}, "a", -1},
		{[]string{"b", "d", "f"}, "b", 0},
		{[]string{"b", "d", "f"}, "c", -1},
		{[]string{"b", "d", "f"}, "d", 1},
		{[]string{"b", "d", "f"}, "e", -1},
		{[]string{"b", "d", "f"}, "f", 2},
		{[]string{"b", "d", "f"}, "g", -1},
		// Single-element slice
		{[]string{"a"}, "a", 0},
		{[]string{"a"}, "b", -1},
		// Slice with duplicates - Search can return any index where 'b' is found.
		{[]string{"a", "b", "b", "c"}, "b", 1}, // Target for test.s[got] == test.v
		// Slice with all same elements
		{[]string{"a", "a", "a", "a"}, "a", 0}, // Target for test.s[got] == test.v
		// Value at beginning
		{[]string{"a", "b", "c", "d", "e"}, "a", 0},
		// Value at end
		{[]string{"a", "b", "c", "d", "e"}, "e", 4},
		// Value in middle
		{[]string{"a", "b", "c", "d", "e"}, "c", 2},
	}

	for i, test := range tests {
		got := Search(test.s, test.v)
		if test.want == -1 { // Case: Element should not be found
			if got != -1 {
				t.Errorf("%d: Search(%v, %q) = %d, want -1 (element should NOT be found)", i, test.s, test.v, got)
			}
		} else { // Case: Element should be found
			if got == -1 {
				t.Errorf("%d: Search(%v, %q) = -1, want an index for %q (element SHOULD be found)", i, test.s, test.v, test.v)
			} else if test.s[got] != test.v {
				t.Errorf("%d: Search(%v, %q) = %d, but s[%d] is %q, not %q (found WRONG element)", i, test.s, test.v, got, got, test.s[got], test.v)
			}
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
		// Slices with mixed duplicates
		{
			[]string{"a", "b", "b", "c"},
			"b", // value with duplicates
			1,   // should point to the first 'b'
		},
		{
			[]string{"a", "b", "b", "c"},
			"a", // value before duplicates
			0,
		},
		{
			[]string{"a", "b", "b", "c"},
			"c", // value after duplicates
			3,
		},
		{
			[]string{"a", "a", "b", "c", "c"},
			"a",
			0,
		},
		{
			[]string{"a", "a", "b", "c", "c"},
			"b",
			2,
		},
		{
			[]string{"a", "a", "b", "c", "c"},
			"c",
			3,
		},
		{
			[]string{"a", "b", "c"}, // no duplicates
			"b",
			1,
		},
	}

	for i, test := range tests {
		if gotInt := BisectLeft(test.s, test.v); gotInt != test.want {
			t.Errorf("%d: BisectLeft(%v, %v) = %d, want %d", i, test.s, test.v, gotInt, test.want)
		}
	}
}

func TestBisectRight(t *testing.T) {
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
			1,
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
			1,
		},
		{
			[]string{"b", "d", "f"},
			"c",
			1,
		},
		{
			[]string{"b", "d", "f"},
			"d",
			2,
		},
		{
			[]string{"b", "d", "f"},
			"e",
			2,
		},
		{
			[]string{"b", "d", "f"},
			"f",
			3,
		},
		{
			[]string{"b", "d", "f"},
			"g",
			3,
		},
		{
			[]string{"b", "b", "b"},
			"b",
			3,
		},
		// Slices with mixed duplicates
		{
			[]string{"a", "b", "b", "c"},
			"b", // value with duplicates
			3,   // should point to the index after the last 'b'
		},
		{
			[]string{"a", "b", "b", "c"},
			"a", // value before duplicates
			1,
		},
		{
			[]string{"a", "b", "b", "c"},
			"c", // value after duplicates
			4,
		},
		{
			[]string{"a", "a", "b", "c", "c"},
			"a",
			2,
		},
		{
			[]string{"a", "a", "b", "c", "c"},
			"b",
			3,
		},
		{
			[]string{"a", "a", "b", "c", "c"},
			"c",
			5,
		},
		{
			[]string{"a", "b", "c"}, // no duplicates
			"b",
			2,
		},
	}

	for i, test := range tests {
		if got := BisectRight(test.s, test.v); got != test.want {
			t.Errorf("%d: BisectRight(%v, %v) = %d, want %d", i, test.s, test.v, got, test.want)
		}
	}
}
