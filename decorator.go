package fun

// Decorator is a generic signature for a decorator function.
type Decorator[T any] func(T) T

// ID is an identity function, a function that returns the same value it receives.
// See: https://en.wikipedia.org/wiki/Identity_function
func ID[T any](t T) T { return t }

// Decorate applies the decorator to the provided value.
func (d Decorator[T]) Decorate(t T) T { return d(t) }

// DecorateAny decorates the provided value with any number of decorators.
func DecorateAny[T any](t T, decorators ...Decorator[T]) T { return Decorate(t, decorators) }

// Decorate decorates the provided value with the provided decorators.
func Decorate[T any](t T, decorators []Decorator[T]) T {
	for _, d := range decorators {
		t = d.Decorate(t)
	}
	return t
}
