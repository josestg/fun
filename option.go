package fun

// Option is a signature for Functional Options of type T.
// See: https://www.josestg.com/posts/design-pattern/functional-options/
type Option[T any] func(*T)

// Apply applies the option to the dst.
func (f Option[T]) Apply(dst *T) { f(dst) }

// ApplyOptions applies the options to the dst.
func ApplyOptions[T any](dst *T, opts []Option[T]) {
	for _, opt := range opts {
		opt.Apply(dst)
	}
}
