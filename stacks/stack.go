// Package stacks defines the Stack interface and all of it's implementations.
package stacks

// Stack is the interface for stacks implementations.
// A Stack is an abstract data type that is a collection of
// elements where insertions and removals happen in LIFO,
// last-in, first-out.
type Stack[T any] interface {
	// Push adds a value to the top of the Stack.
	Push(T)
	// Pop attempts to remove and return the value at the top
	// of the Stack and reports whether it succeeded.
	Pop() (T, bool)
	// Peek attempts to return the value at the top of the Stack
	// without removing it and reports whether it succeeded.
	Peek() (T, bool)
	// Len returns Stack's length.
	Len() int
}
