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
	w, err := q.Projection("foo")
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}
	actual := fmt.Sprintf("%v", slices.Collect(w.List()))
	expect := fmt.Sprintf("%v", []map[string]any{foo})
	if actual != expect {
		t.Errorf("expected %v got %v", expect, actual)
	}

	t.Run("keep relation with schema subset", func(t *testing.T) {
		foobarbaz := map[string]any{"foo": 1, "bar": 2.0, "baz": "3"}
		q := newTable(t, schema.FromTuple(foobarbaz), foo, bar, foobarbaz)
		w, err := q.Projection("foo", "bar")
		if err != nil {
			t.Fatalf("unexpected error %v", err)
		}
		actual := fmt.Sprintf("%v", slices.Collect(w.List()))
		expect := fmt.Sprintf("%v", []map[string]any{foo, bar, foobar})
		if actual != expect {
			t.Errorf("expected %v got %v", expect, actual)
		}
	})

	t.Run("mismatch", func(t *testing.T) {
		_, err := q.Projection("qux")
		if err != schema.ErrMismatch {
			t.Errorf("expected error %v got %v", schema.ErrMismatch, err)
		}
	})
}
