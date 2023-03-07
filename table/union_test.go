package table_test

import (
	"fmt"
	"slices"
	"testing"

	"github.com/martindrlik/rex/schema"
	"github.com/martindrlik/rex/tuple"
)

func TestUnion(t *testing.T) {
	foo := map[string]any{"foo": 1}
	bar := map[string]any{"bar": 1.0}
	bar2 := map[string]any{"bar": 2.0}
	foobar := tuple.Merge(foo, bar)
	foobar2 := tuple.Merge(foo, bar2)

	q := newTable(t, schema.FromTuple(foobar), foo, foobar)
	r := newTable(t, schema.FromTuple(foobar), bar, foobar2)

	w, err := q.Union(r)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	actual := fmt.Sprintf("%v", slices.Collect(w.List()))
	expect := fmt.Sprintf("%v", []map[string]any{foo, foobar, foobar2, bar})
	if actual != expect {
		t.Errorf("expected %v got %v", expect, actual)
	}

	t.Run("mismatch", func(t *testing.T) {
		baz := map[string]any{"baz": "qux"}
		r := newTable(t, schema.FromTuple(baz), baz)
		_, err := q.Union(r)
		if err != schema.ErrMismatch {
			t.Errorf("expected %v got %v", schema.ErrMismatch, err)
		}
	})
}
