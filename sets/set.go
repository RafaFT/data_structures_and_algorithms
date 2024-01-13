// Package sets defines the Set interface and all of it's implementations.
// A Set is defined as an abstract data structure that prevents duplicate values.
package sets

// Set is the interface for set implementations.
// A Set is an abstract data structure that prevents duplicate values.
type Set[T comparable] interface {
	// Add attempts to insert value to Set and reports whether it succeed.
	Add(value T) bool
	// Has reports whether value happens in Set.
	Has(value T) bool
	// Remove removes value from Set and reports whether it was found.
	Remove(value T) bool
}
