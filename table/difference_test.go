package table_test

import (
	"fmt"
	"slices"
	"testing"

	"github.com/martindrlik/rex/schema"
	"github.com/martindrlik/rex/tuple"
)

func TestDifference(t *testing.T) {
	foo1 := map[string]any{"foo": 1}
	foo2 := map[string]any{"foo": 2}
	bar := map[string]any{"bar": 3.0}

	foo1bar := tuple.Merge(foo1, bar)
	foo2bar := tuple.Merge(foo2, bar)

	q := newTable(t, schema.FromTuple(foo1bar), foo1, foo2, foo1bar, foo2bar)
	r := newTable(t, schema.FromTuple(foo1bar), foo1, foo1bar)

	w, err := q.Difference(r)
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}

	actual := fmt.Sprintf("%v", slices.Collect(w.List()))
	expect := fmt.Sprintf("%v", []map[string]any{foo1, foo2, foo2bar})
	if actual != expect {
		t.Errorf("expected %v got %v", expect, actual)
	}

	t.Run("mismatch", func(t *testing.T) {
		baz := map[string]any{"baz": "qux"}
		r := newTable(t, schema.FromTuple(baz), baz)
		_, err := q.Difference(r)
		if err != schema.ErrMismatch {
			t.Errorf("expected error %v got %v", schema.ErrMismatch, err)
		}
	})
}
