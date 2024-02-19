package fun

import "testing"

func TestDecorateAny(t *testing.T) {
	t.Run("scalar", func(t *testing.T) {
		inc := func(x int) int { return x + 1 }
		dec := func(x int) int { return x - 1 }
		dbl := func(x int) int { return 2 * x }
		n := DecorateAny(1, inc, inc, dbl, dec)
		if n != 5 {
			t.Errorf("expected 5, got %d", n)
		}
	})

	t.Run("function", func(t *testing.T) {
		one := func() int { return 1 }
		inc := func(f func() int) func() int { return func() int { return f() + 1 } }
		dec := func(f func() int) func() int { return func() int { return f() - 1 } }
		dbl := func(f func() int) func() int { return func() int { return 2 * f() } }

		f := DecorateAny(one, inc, inc, dbl, dec)
		n := f()
		if n != 5 {
			t.Errorf("expected 5, got %d", n)
		}
	})
}

func TestID(t *testing.T) {
	n := ID(1)
	if n != 1 {
		t.Errorf("expected 1, got %d", n)
	}

	f := ID(func() int { return 1 })
	n = f()
	if n != 1 {
		t.Errorf("expected 1, got %d", n)
	}
}
