package table_test

import (
	"fmt"
	"slices"
	"testing"

	"github.com/martindrlik/rex/schema"
	"github.com/martindrlik/rex/tuple"
)

func TestProjection(t *testing.T) {
	foo := map[string]any{"foo": 1}
	bar := map[string]any{"bar": 2.0}
	foobar := tuple.Merge(foo, bar)
	q := newTable(t, schema.FromTuple(foobar), foo, bar, foobar)
	w, ok := q.Projection("foo")
	if !ok {
		t.Fatal("unexpected empty result")
	}
	actual := fmt.Sprintf("%v", slices.Collect(w.List()))
	expect := fmt.Sprintf("%v", []map[string]any{foo})
	if actual != expect {
		t.Errorf("expected %v got %v", expect, actual)
	}

	t.Run("empty result", func(t *testing.T) {
		_, ok := q.Projection("qux")
		if ok {
			t.Error("unexpected non-empty result")
		}
	})
}
