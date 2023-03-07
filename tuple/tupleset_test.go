package tuple_test

import (
	"fmt"
	"testing"

	"github.com/martindrlik/rex/tuple"
)

func TestTupleSet(t *testing.T) {
	foobar := map[string]any{"foo": 1, "bar": 2.0}
	ts := tuple.TupleSet{}
	ts.Add(foobar)
	actual := fmt.Sprintf("%v", ts)
	expect := "[map[bar:2 foo:1]]"
	if actual != expect {
		t.Errorf("expected %v got %v", expect, actual)
	}
	t.Run("has", func(t *testing.T) {
		barbaz := map[string]any{"bar": 2.0, "baz": "3"}
		if !ts.Has(foobar) {
			t.Errorf("expected %v to contain %v", ts, foobar)
		}
		if ts.Has(barbaz) {
			t.Errorf("expected %v to not contain %v", ts, barbaz)
		}
	})
}

func TestDelete(t *testing.T) {
	foobar := map[string]any{"foo": 1, "bar": 2.0}
	ts := tuple.TupleSet{}
	ts.Add(foobar)
	ts.Delete(foobar)
	if ts.Has(foobar) {
		t.Errorf("expected %v to not contain %v", ts, foobar)
	}
}
