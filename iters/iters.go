package iters

import (
	"github.com/josestg/fun/prd"
)

// Some returns true if at least one element in the slice satisfies the predicate.
func Some[S ~[]E, E any](s S, p prd.Predicate[E]) bool {
	for _, v := range s {
		if p(v) {
			return true
		}
	}
	return false
}

// All returns true if all elements in the slice satisfy the predicate.
func All[S ~[]E, E any](s S, p prd.Predicate[E]) bool {
	for _, v := range s {
		if !p(v) {
			return false
		}
	}
	return true
}

// Filter returns a new slice containing only the elements that satisfy the predicate.
func Filter[S ~[]E, E any](s S, p prd.Predicate[E]) S {
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
func Count[S ~[]E, E any](s S, p prd.Predicate[E]) int {
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
func Find[S ~[]E, E any](s S, p prd.Predicate[E]) (E, bool) {
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

// Equal returns true if the two slices are equal.
// Two slices are equal if they have the same length and all elements are equal.
func Equal[S ~[]E, E comparable](a, b S) bool {
	return EqualBy(a, b, func(e1 E, e2 E) bool { return e1 == e2 })
}

// EqualBy returns true if the two slice are satisfied the quality function.
// Two slices are equal if they have the same length and all elements are satisfied the quality function.
func EqualBy[S1 ~[]E1, S2 ~[]E2, E1, E2 any](a S1, b S2, eq func(E1, E2) bool) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if !eq(a[i], b[i]) {
			return false
		}
	}

	return true
}

// Index returns the index of the first occurrence of the value in the slice.
func Index[S ~[]E, E comparable](s S, v E) int {
	return IndexBy(s, prd.Eq(v))
}

// LastIndex returns the index of the last occurrence of the value in the slice.
func LastIndex[S ~[]E, E comparable](s S, v E) int {
	return LastIndexBy(s, prd.Eq(v))
}

// IndexBy returns the index of the first element in the slice that satisfies the predicate.
func IndexBy[S ~[]E, E any](s S, p prd.Predicate[E]) int {
	for i := range s {
		if p(s[i]) {
			return i
		}
	}
	return -1
}

// LastIndexBy returns the index of the last element in the slice that satisfies the predicate.
func LastIndexBy[S ~[]E, E any](s S, p prd.Predicate[E]) int {
	for i := len(s) - 1; i >= 0; i-- {
		if p(s[i]) {
			return i
		}
	}
	return -1
}

// Contains returns true if the value is in the slice.
func Contains[S ~[]E, E comparable](s S, v E) bool {
	return ContainsBy(s, prd.Eq(v))
}

// ContainsBy returns true if at least one element in the slice satisfies the predicate.
func ContainsBy[S ~[]E, E comparable](s S, p prd.Predicate[E]) bool {
	return IndexBy(s, p) != -1
}
