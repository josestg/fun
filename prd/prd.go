package prd

import "cmp"

// Predicate checks if a value satisfies a condition.
type Predicate[T any] func(T) bool

// Neq checks if a value is not equal to another.
func Neq[T comparable](v T) Predicate[T] { return func(x T) bool { return x != v } }

// Eq checks if a value is equal to another.
func Eq[T comparable](v T) Predicate[T] { return func(x T) bool { return x == v } }

// Lt checks if a value is less than another.
func Lt[T cmp.Ordered](v T) Predicate[T] { return func(x T) bool { return x < v } }

// Gt checks if a value is greater than another.
func Gt[T cmp.Ordered](v T) Predicate[T] { return func(x T) bool { return x > v } }

// Lte checks if a value is less than or equal to another.
func Lte[T cmp.Ordered](v T) Predicate[T] { return func(x T) bool { return x <= v } }

// Gte checks if a value is greater than or equal to another.
func Gte[T cmp.Ordered](v T) Predicate[T] { return func(x T) bool { return x >= v } }

// Empty checks if a value is the zero value.
func Empty[T comparable]() Predicate[T] {
	var zero T
	return Eq(zero)
}

// InBetween checks if a value is in between of min and max values (inclusive).
func InBetween[T cmp.Ordered](minValue, maxValue T) Predicate[T] {
	return And(Gte(minValue), Lte(maxValue))
}

// ExBetween checks if a value is in between of min and max values (exclusive).
func ExBetween[T cmp.Ordered](minValue, maxValue T) Predicate[T] {
	return And(Gt(minValue), Lt(maxValue))
}

// Or returns a predicate that is true if either of the given predicates is true.
func Or[T any](p, q Predicate[T]) Predicate[T] { return func(x T) bool { return p(x) || q(x) } }

// And returns a predicate that is true if both of the given predicates are true.
func And[T any](p, q Predicate[T]) Predicate[T] { return func(x T) bool { return p(x) && q(x) } }

// Not returns a predicate that is true if the given predicate is false.
func Not[T any](p Predicate[T]) Predicate[T] { return func(x T) bool { return !p(x) } }
