package fun

import (
	"testing"

	"github.com/josestg/fun/cmp"
)

func TestFirst(t *testing.T) {
	n := First(cmp.Not(cmp.Empty[int]()), 0, 0, 0, 4, 5)
	if n != 4 {
		t.Errorf("expect 4 but got %d", n)
	}

	n = First(cmp.Not(cmp.Empty[int]()), 1, 0, 0, 4, 5)
	if n != 1 {
		t.Errorf("expect 1 but got %d", n)
	}

	n = First(cmp.Not(cmp.Empty[int]()), 0, 0, 0)
	if n != 0 {
		t.Errorf("expect 0 but got %d", n)
	}
}
