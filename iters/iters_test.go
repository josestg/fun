package iters

import (
	"testing"

	"github.com/josestg/fun/prd"
)

type Task struct {
	ID        int
	Title     string
	Completed bool
}

func Completed(t Task) bool { return t.Completed }

func Complete(t Task) Task {
	t.Completed = true
	return t
}

var tasks = []Task{
	{1, "one", true},
	{2, "two", false},
	{3, "three", true},
}

func TestSome(t *testing.T) {
	s := []int{1, 2, 3, 4, 5}
	if !Some(s, prd.Eq(3)) {
		t.Error("expect found some value that equal to 3")
	}

	if Some(s, prd.Gt(5)) {
		t.Error("expect no value that greater than 5")
	}

	ok := Some(tasks, Completed)
	if !ok {
		t.Error("expect some task is completed")
	}
}

func TestAll(t *testing.T) {
	s := []int{1, 2, 3, 4, 5}
	if !All(s, prd.Gte(1)) {
		t.Error("expect all values greater or equal to 1")
	}

	if All(s, prd.Gt(1)) {
		t.Error("not all values is greater or equal to 1")
	}

	ok := All(tasks, Completed)
	if ok {
		t.Error("some task are not completed")
	}
}

func TestFilter(t *testing.T) {
	completedTasks := Filter(tasks, Completed)
	ok := All(completedTasks, Completed)
	if !ok {
		t.Error("expect all task are completed")
	}
}

func TestMap(t *testing.T) {
	ok := All(Map(tasks, Complete), Completed)
	if !ok {
		t.Error("expect all task are completed")
	}
}

func TestCount(t *testing.T) {
	s := []int{1, 2, 3, 4, 5}
	n := Count(s, prd.Lt(3))
	if n != 2 {
		t.Error("expect there are 2 items < 3")
	}
}

func TestFind(t *testing.T) {
	s := []int{1, 2, 3, 4, 5}
	v, ok := Find(s, prd.Gt(3))
	if !ok {
		t.Error("expect found value > 3")
	}
	if v != 4 {
		t.Errorf("expect found value 4, got %d", v)
	}

	v, ok = Find(s, prd.Gt(5))
	if ok {
		t.Error("expect not found value > 5")
	}
}

func TestEach(t *testing.T) {
	s := []int{1, 2, 3, 4, 5}
	Each(s, func(v int) {
		if v > 5 {
			t.Errorf("expect value <= 5, got %d", v)
		}
	})
}

func TestEqual(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	b := []int{1, 2, 3, 4, 5}
	if !Equal(a, b) {
		t.Error("expect equal")
	}

	c := []int{1, 2, 3, 4}
	if Equal(a, c) {
		t.Error("expect not equal")
	}

	d := []int{1, 2, 3, 4, 6}
	if Equal(a, d) {
		t.Error("expect not equal")
	}
}

func TestEqualOf(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	b := []int{5, 6, 7, 8, 9}
	if !EqualOf(a, b, func(e1 int, e2 int) bool {
		return e1 < 10 && e2 < 10
	}) {
		t.Error("expect equal")
	}
}

func TestIndex(t *testing.T) {
	s := []int{1, 2, 1}
	i := Index(s, 1)
	if i != 0 {
		t.Errorf("expect index 0, got %d", i)
	}

	i = Index(s, 3)
	if i != -1 {
		t.Errorf("expect index -1, got %d", i)
	}
}

func TestLastIndex(t *testing.T) {
	s := []int{1, 2, 1}
	i := LastIndex(s, 1)
	if i != 2 {
		t.Errorf("expect index 2, got %d", i)
	}

	i = LastIndex(s, 3)
	if i != -1 {
		t.Errorf("expect index -1, got %d", i)
	}
}

func TestIndexOf(t *testing.T) {
	i := IndexOf(tasks, Completed)
	if i != 0 {
		t.Errorf("expect index 0, got %d", i)
	}
}

func TestLastIndexOf(t *testing.T) {
	i := LastIndexOf(tasks, Completed)
	if i != 2 {
		t.Errorf("expect index 2, got %d", i)
	}
}

func TestContains(t *testing.T) {
	s := []int{1, 2, 3, 4, 5}
	if !Contains(s, 3) {
		t.Error("expect contains 3")
	}

	if Contains(s, 6) {
		t.Error("expect not contains 6")
	}
}
