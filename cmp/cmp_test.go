package cmp

import "testing"

func TestEq(t *testing.T) {
	one := Eq(1)
	expectTrue(t, one(1))
	expectFalse(t, one(2))
}

func TestNeq(t *testing.T) {
	one := Neq(1)
	expectFalse(t, one(1))
	expectTrue(t, one(2))
}

func TestGt(t *testing.T) {
	one := Gt(1)
	expectFalse(t, one(1))
	expectTrue(t, one(2))
}

func TestGte(t *testing.T) {
	one := Gte(1)
	expectTrue(t, one(1))
	expectTrue(t, one(2))
	expectFalse(t, one(0))
}

func TestLt(t *testing.T) {
	one := Lt(1)
	expectFalse(t, one(1))
	expectTrue(t, one(0))
}

func TestLte(t *testing.T) {
	one := Lte(1)
	expectTrue(t, one(1))
	expectTrue(t, one(0))
	expectFalse(t, one(2))
}

func TestEmpty(t *testing.T) {
	empty := Empty[int]()
	expectTrue(t, empty(0))
	expectFalse(t, empty(1))
}

func TestInBetween(t *testing.T) {
	inBetween := InBetween(1, 3)
	expectTrue(t, inBetween(1))
	expectTrue(t, inBetween(2))
	expectTrue(t, inBetween(3))
	expectFalse(t, inBetween(0))
	expectFalse(t, inBetween(4))
}

func TestExBetween(t *testing.T) {
	exBetween := ExBetween(1, 3)
	expectFalse(t, exBetween(1))
	expectTrue(t, exBetween(2))
	expectFalse(t, exBetween(3))
	expectFalse(t, exBetween(0))
	expectFalse(t, exBetween(4))
}

func TestOr(t *testing.T) {
	oneOrTwo := Or(Eq(1), Eq(2))
	expectTrue(t, oneOrTwo(1))
	expectTrue(t, oneOrTwo(2))
	expectFalse(t, oneOrTwo(3))
}

func TestAnd(t *testing.T) {
	oneAndTwo := And(Eq(1), Eq(2))
	expectFalse(t, oneAndTwo(1))
	expectFalse(t, oneAndTwo(2))
	expectFalse(t, oneAndTwo(3))
}

func expectTrue(t *testing.T, b bool) {
	if !b {
		t.Error("expect true")
	}
}

func expectFalse(t *testing.T, b bool) {
	if b {
		t.Error("expect false")
	}
}
