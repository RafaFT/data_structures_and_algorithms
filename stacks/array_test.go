package stacks

import (
	"fmt"
	"reflect"
	"testing"
)

func TestStackArray(t *testing.T) {
	t.Run("Push", func(t *testing.T) {
		tests := []struct {
			stack *StackArray[int]
			value int
			want  *StackArray[int]
		}{
			{
				&StackArray[int]{},
				2,
				&StackArray[int]{[]int{2}},
			},
			{
				&StackArray[int]{[]int{2}},
				1,
				&StackArray[int]{[]int{2, 1}},
			},
		}

		for i, test := range tests {
			before := fmt.Sprint(test.stack)
			test.stack.Push(test.value)

			if !reflect.DeepEqual(test.stack, test.want) {
				t.Errorf("%d: %v.Push(%v) != %v", i, before, test.value, test.want)
			}
		}
	})

	t.Run("Pop", func(t *testing.T) {
		tests := []struct {
			stack     *StackArray[int]
			wantValue int
			wantBool  bool
			wantStack *StackArray[int]
		}{
			{
				&StackArray[int]{[]int{2, 1}},
				1,
				true,
				&StackArray[int]{[]int{2}},
			},
			{
				&StackArray[int]{[]int{2}},
				2,
				true,
				&StackArray[int]{[]int{}},
			},
			{
				&StackArray[int]{},
				0,
				false,
				&StackArray[int]{},
			},
		}

		for i, test := range tests {
			before := fmt.Sprint(test.stack)

			if gotValue, gotBool := test.stack.Pop(); gotValue != test.wantValue ||
				gotBool != test.wantBool ||
				!reflect.DeepEqual(test.stack, test.wantStack) {
				t.Errorf(
					"%d: %v.Pop() != (%v, %v), want (%v, %v) and %v",
					i, before, gotValue, gotBool, test.wantValue, test.wantBool, test.wantStack,
				)
			}
		}
	})

	t.Run("Peek", func(t *testing.T) {
		tests := []struct {
			stack     *StackArray[int]
			wantValue int
			wantBool  bool
			wantStack *StackArray[int]
		}{
			{
				&StackArray[int]{},
				0,
				false,
				&StackArray[int]{},
			},
			{
				&StackArray[int]{[]int{2}},
				2,
				true,
				&StackArray[int]{[]int{2}},
			},
			{
				&StackArray[int]{[]int{2, 1}},
				1,
				true,
				&StackArray[int]{[]int{2, 1}},
			},
		}

		for i, test := range tests {
			if gotValue, gotBool := test.stack.Peek(); gotValue != test.wantValue ||
				gotBool != test.wantBool ||
				!reflect.DeepEqual(test.stack, test.wantStack) {
				t.Errorf(
					"%d: %v.Peek() != (%v, %v), want (%v, %v) and %v",
					i, test.stack, gotValue, gotBool, test.wantValue, test.wantBool, test.wantStack,
				)
			}
		}
	})
}
