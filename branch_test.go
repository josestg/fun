package fun

import "testing"

func TestIf(t *testing.T) {
	n := If(true, 1, 2)
	if n != 1 {
		t.Errorf("expected 1, got %d", n)
	}

	n = If(false, 1, 2)
	if n != 2 {
		t.Errorf("expected 2, got %d", n)
	}
}

func TestUnless(t *testing.T) {
	n := Unless(false, 1, 2)
	if n != 1 {
		t.Errorf("expected 1, got %d", n)
	}

	n = Unless(true, 1, 2)
	if n != 2 {
		t.Errorf("expected 2, got %d", n)
	}
}
