package lists

import (
	"testing"
)

func panics(fn func()) (panicked bool) {
	defer func() {
		if e := recover(); e != nil {
			panicked = true
		}
	}()

	fn()

	return panicked
}

func equals[T comparable](l1, l2 LinkedList[T]) bool {
	if l1.len != l2.len {
		return false
	}

	if l1.len == 0 && (l1.head != nil || l2.head != nil) {
		return false
	}

	n1, n2 := l1.head, l2.head
	for i := 0; i < l1.len; i++ {
		if n1.value != n2.value {
			return false
		}

		n1, n2 = n1.next, n2.next
	}

	return true
}

func TestOutOfBounds(t *testing.T) {
	tests := []struct {
		ll    LinkedList[string]
		index int
	}{
		{
			LinkedList[string]{
				len:  0,
				head: nil,
			},
			-1,
		},
		{
			LinkedList[string]{
				len:  0,
				head: nil,
			},
			0,
		},
		{
			LinkedList[string]{
				len: 1,
				head: &node[string]{
					value: "",
					next:  nil,
				},
			},
			1,
		},
		{
			LinkedList[string]{
				len: 1,
				head: &node[string]{
					value: "",
					next:  nil,
				},
			},
			2,
		},
	}

	t.Run("Read", func(t *testing.T) {
		for i, test := range tests {
			if panicked := panics(func() { test.ll.Read(test.index) }); !panicked {
				t.Errorf("%d: LinkedList.Read(%d) expected to panic", i, test.index)
			}
		}
	})

	t.Run("Insert", func(t *testing.T) {
		for i, test := range tests {
			index := test.index
			if index == test.ll.len {
				index++
			}

			if panicked := panics(func() { test.ll.Insert("", index) }); !panicked {
				t.Errorf("%d: LinkedList.Insert(%d) expected to panicked", i, index)
			}
		}
	})

	t.Run("Delete", func(t *testing.T) {
		for i, test := range tests {
			if panicked := panics(func() { test.ll.Delete(test.index) }); !panicked {
				t.Errorf("%d: LinkedList.Delete(%d) expected to panic", i, test.index)
			}
		}
	})
}

func Test(t *testing.T) {
	l := LinkedList[string]{
		head: &node[string]{
			value: "a",
			next: &node[string]{
				value: "b",
				next: &node[string]{
					value: "c",
					next:  nil,
				},
			},
		},
		len: 3,
	}

	t.Run("Read", func(t *testing.T) {
		tests := []struct {
			index int
			want  string
		}{
			{
				0,
				"a",
			},
			{
				1,
				"b",
			},
			{
				2,
				"c",
			},
		}

		for i, test := range tests {
			if got := l.Read(test.index); got != test.want {
				t.Errorf("%d: LinkedList.Read(%d) = %s, got %s", i, test.index, test.want, got)
			}
		}
	})

	t.Run("Search", func(t *testing.T) {
		tests := []struct {
			value string
			want  int
		}{
			{
				"",
				-1,
			},
			{
				"a",
				0,
			},
			{
				"b",
				1,
			},
			{
				"c",
				2,
			},
			{
				"d",
				-1,
			},
		}

		for i, test := range tests {
			if got := l.Search(test.value); got != test.want {
				t.Errorf("%d: LinkedList.Search(%s) = %d, got %d", i, test.value, test.want, got)
			}
		}
	})
}

func TestInsert(t *testing.T) {
	tests := []struct {
		ll    LinkedList[string]
		value string
		index int
		want  LinkedList[string]
	}{
		{
			LinkedList[string]{
				0,
				nil,
			},
			"b",
			0,
			LinkedList[string]{
				1,
				&node[string]{
					value: "b",
					next:  nil,
				},
			},
		},
		{
			LinkedList[string]{
				1,
				&node[string]{
					value: "b",
					next:  nil,
				},
			},
			"a",
			0,
			LinkedList[string]{
				2,
				&node[string]{
					value: "a",
					next: &node[string]{
						value: "b",
						next:  nil,
					},
				},
			},
		},
		{
			LinkedList[string]{
				2,
				&node[string]{
					value: "a",
					next: &node[string]{
						value: "b",
						next:  nil,
					},
				},
			},
			"d",
			2,
			LinkedList[string]{
				3,
				&node[string]{
					value: "a",
					next: &node[string]{
						value: "b",
						next: &node[string]{
							value: "d",
							next:  nil,
						},
					},
				},
			},
		},
		{
			LinkedList[string]{
				3,
				&node[string]{
					value: "a",
					next: &node[string]{
						value: "b",
						next: &node[string]{
							value: "d",
							next:  nil,
						},
					},
				},
			},
			"c",
			2,
			LinkedList[string]{
				4,
				&node[string]{
					value: "a",
					next: &node[string]{
						value: "b",
						next: &node[string]{
							value: "c",
							next: &node[string]{
								value: "d",
								next:  nil,
							},
						},
					},
				},
			},
		},
	}

	for i, test := range tests {
		test.ll.Insert(test.value, test.index)

		if !equals(test.ll, test.want) {
			t.Errorf("%d: LinkedList.Insert(%s, %d)", i, test.value, test.index)
		}
	}
}

func TestDelete(t *testing.T) {
	tests := []struct {
		ll    LinkedList[string]
		index int
		want  LinkedList[string]
	}{
		{
			LinkedList[string]{
				1,
				&node[string]{
					value: "a",
					next:  nil,
				},
			},
			0,
			LinkedList[string]{
				0,
				nil,
			},
		},
		{
			LinkedList[string]{
				2,
				&node[string]{
					value: "a",
					next: &node[string]{
						value: "b",
						next:  nil,
					},
				},
			},
			0,
			LinkedList[string]{
				1,
				&node[string]{
					value: "b",
					next:  nil,
				},
			},
		},
		{
			LinkedList[string]{
				2,
				&node[string]{
					value: "a",
					next: &node[string]{
						value: "b",
						next:  nil,
					},
				},
			},
			1,
			LinkedList[string]{
				1,
				&node[string]{
					value: "a",
					next:  nil,
				},
			},
		},
	}

	for i, test := range tests {
		test.ll.Delete(test.index)

		if !equals(test.ll, test.want) {
			t.Errorf("%d: LinkedList.Delete(%d)", i, test.index)
		}
	}
}
