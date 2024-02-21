package slices

import (
	"github.com/josestg/fun/cmp"
)

// Some returns true if at least one element in the slice satisfies the predicate.
func Some[T any](s []T, p cmp.Predicate[T]) bool {
	for _, v := range s {
		if p(v) {
			return true
		}
	}
	return false
}

// All returns true if all elements in the slice satisfy the predicate.
func All[T any](s []T, p cmp.Predicate[T]) bool {
	for _, v := range s {
		if !p(v) {
			return false
		}
	}
	return true
}

// Filter returns a new slice containing only the elements that satisfy the predicate.
func Filter[T any](s []T, p cmp.Predicate[T]) []T {
	n := Count(s, p)
	return Fold(
		s,
		make([]T, 0, n),
		func(xs []T, x T) []T {
			if p(x) {
				xs = append(xs, x)
			}
			return xs
		},
	)
}

// Count returns the number of elements in the slice that satisfy the predicate.
func Count[T any](s []T, p cmp.Predicate[T]) int {
	return Fold(s, 0, func(n int, x T) int {
		if p(x) {
			n++
		}
		return n
	})
}

// Map returns a new slice containing the result of applying the function to each element in the slice.
func Map[T, U any](s []T, f func(T) U) []U {
	return Fold(
		s,
		make([]U, 0, len(s)),
		func(xs []U, x T) []U { return append(xs, f(x)) },
	)
}

// Find returns the first element in the slice that satisfies the predicate.
func Find[T any](s []T, p cmp.Predicate[T]) (T, bool) {
	for _, v := range s {
		if p(v) {
			return v, true
		}
	}
	var zero T
	return zero, false
}

// Fold applies the function to each element in the slice, accumulating the result.
func Fold[T, U any](s []T, z U, f func(U, T) U) U {
	r := z
	for _, v := range s {
		r = f(r, v)
	}
	return r
}

// Each applies the function to each element in the slice.
func Each[T any](s []T, f func(T)) {
	for _, v := range s {
		f(v)
	}
}
