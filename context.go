package fun

import (
	"context"
	"time"
)

// pair is a generic pair type.
type pair[F, S any] struct {
	f F
	s S
}

// pairOf returns a new pair.
func pairOf[F, S any](f F, s S) pair[F, S] {
	return pair[F, S]{f: f, s: s}
}

// RunWithContext runs the function f with the given context.
func RunWithContext[T any](ctx context.Context, f func() (T, error)) (ret T, err error) {
	// Check if the context is already done, to prevent unnecessary work.
	select {
	case <-ctx.Done():
		return ret, ctx.Err()
	default:
	}

	res := make(chan pair[T, error], 1)
	go func() {
		res <- pairOf(f())
		close(res)
	}()

	select {
	case <-ctx.Done():
		return ret, ctx.Err()
	case r := <-res:
		return r.f, r.s
	}
}

// RunWithTimeout runs the function f with the given timeout.
func RunWithTimeout[T any](timeout time.Duration, f func() (T, error)) (T, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	return RunWithContext(ctx, f)
}

// RunWithDeadline runs the function f with the given deadline.
func RunWithDeadline[T any](deadline time.Time, f func() (T, error)) (T, error) {
	ctx, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel()
	return RunWithContext(ctx, f)
}
