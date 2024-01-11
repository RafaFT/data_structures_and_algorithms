package sets

type Set[T comparable] interface {
	Add(T) bool
	Has(T) bool
	Remove(T) bool
}
