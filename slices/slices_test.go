package slices

import (
	"testing"

	"github.com/josestg/fun/cmp"
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
	if !Some(s, cmp.Eq(3)) {
		t.Error("expect found some value that equal to 3")
	}

	if Some(s, cmp.Gt(5)) {
		t.Error("expect no value that greater than 5")
	}

	ok := Some(tasks, Completed)
	if !ok {
		t.Error("expect some task is completed")
	}
}

func TestAll(t *testing.T) {
	s := []int{1, 2, 3, 4, 5}
	if !All(s, cmp.Gte(1)) {
		t.Error("expect all values greater or equal to 1")
	}

	if All(s, cmp.Gt(1)) {
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
	n := Count(s, cmp.Lt(3))
	if n != 2 {
		t.Error("expect there are 2 items < 3")
	}
}

func TestFind(t *testing.T) {
	s := []int{1, 2, 3, 4, 5}
	v, ok := Find(s, cmp.Gt(3))
	if !ok {
		t.Error("expect found value > 3")
	}
	if v != 4 {
		t.Errorf("expect found value 4, got %d", v)
	}

	v, ok = Find(s, cmp.Gt(5))
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
