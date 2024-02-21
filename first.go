package fun

import "github.com/josestg/fun/prd"

// First returns the first value that satisfies the predicate.
func First[T any](p prd.Predicate[T], x T, xs ...T) T {
	if p(x) {
		return x
	}
	for i := range xs {
		if p(xs[i]) {
			return xs[i]
		}
	}
	return x
}
