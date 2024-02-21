package slices

import (
	"github.com/josestg/fun/cmp"
)

// Some returns true if at least one element in the slice satisfies the predicate.
func Some[S ~[]E, E any](s S, p cmp.Predicate[E]) bool {
	for _, v := range s {
		if p(v) {
			return true
		}
	}
	return false
}

// All returns true if all elements in the slice satisfy the predicate.
func All[S ~[]E, E any](s S, p cmp.Predicate[E]) bool {
	for _, v := range s {
		if !p(v) {
			return false
		}
	}
	return true
}

// Filter returns a new slice containing only the elements that satisfy the predicate.
func Filter[S ~[]E, E any](s S, p cmp.Predicate[E]) S {
	n := Count(s, p)
	return Fold(
		s,
		make(S, 0, n),
		func(xs S, x E) S {
			if p(x) {
				xs = append(xs, x)
			}
			return xs
		},
	)
}

// Count returns the number of elements in the slice that satisfy the predicate.
func Count[S ~[]E, E any](s S, p cmp.Predicate[E]) int {
	return Fold(s, 0, func(n int, x E) int {
		if p(x) {
			n++
		}
		return n
	})
}

// Map returns a new slice containing the result of applying the function to each element in the slice.
func Map[SE ~[]E, ST []T, E, T any](s SE, f func(E) T) ST {
	return Fold(
		s,
		make(ST, 0, len(s)),
		func(xs ST, x E) ST { return append(xs, f(x)) },
	)
}

// Find returns the first element in the slice that satisfies the predicate.
func Find[S ~[]E, E any](s S, p cmp.Predicate[E]) (E, bool) {
	for _, v := range s {
		if p(v) {
			return v, true
		}
	}
	var zero E
	return zero, false
}

// Fold applies the function to each element in the slice, accumulating the result.
func Fold[S ~[]E, E, T any](s S, z T, f func(T, E) T) T {
	r := z
	for _, v := range s {
		r = f(r, v)
	}
	return r
}

// Each applies the function to each element in the slice.
func Each[S ~[]E, E any](s S, f func(E)) {
	for _, v := range s {
		f(v)
	}
}
