package fun

// If returns `then` if `cond` is true, otherwise it returns `els`.
func If[T any](cond bool, then, els T) T {
	if cond {
		return then
	}
	return els
}

// Unless returns `then` if `cond` is false, otherwise it returns `els`.
// It is equivalent to `If(!cond, then, els)`.
func Unless[T any](cond bool, then, els T) T {
	return If(!cond, then, els)
}
