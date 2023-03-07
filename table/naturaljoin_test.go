package table_test

import (
	"fmt"
	"slices"
	"testing"

	"github.com/martindrlik/rex/schema"
	"github.com/martindrlik/rex/tuple"
)

func TestNaturalJoin(t *testing.T) {
	foo := map[string]any{"foo": 1}
	bar := map[string]any{"bar": 2.0}
	baz := map[string]any{"baz": "3"}
	foobar := tuple.Merge(foo, bar)
	foobaz := tuple.Merge(foo, baz)
	q := newTable(t, schema.FromTuple(foobar), foo, foobar)
	r := newTable(t, schema.FromTuple(foobaz), foo, foobaz)
	w := q.NaturalJoin(r)
	actual := fmt.Sprintf("%v", slices.Collect(w.List()))
	expect := fmt.Sprintf("%v", []map[string]any{
		foo,
		foobaz,
		tuple.Merge(foobar, foo),
		tuple.Merge(foobar, foobaz),
	})
	if actual != expect {
		t.Errorf("expected %v got %v", expect, actual)
	}

	t.Run("empty", func(t *testing.T) {
		foobar := map[string]any{"foo": 1, "bar": 2.0}
		baz := map[string]any{"baz": "3"}
		q := newTable(t, schema.FromTuple(foobar))
		r := newTable(t, schema.FromTuple(baz), baz)
		w := q.NaturalJoin(r)
		for range w.List() {
			t.Error("unexpected non-empty result")
		}

	})
}

func TestNaturalJoinCartesianProduct(t *testing.T) {
	foo := map[string]any{"foo": 1}
	bar := map[string]any{"bar": 2.0}
	baz := map[string]any{"baz": "3"}
	pub := map[string]any{"pub": complex(4.0, 1.0)}
	foobar := tuple.Merge(foo, bar)
	bazpub := tuple.Merge(baz, pub)
	q := newTable(t, schema.FromTuple(foobar), foo, foobar)
	r := newTable(t, schema.FromTuple(bazpub), baz, bazpub)
	w := q.NaturalJoin(r)
	actual := fmt.Sprintf("%v", slices.Collect(w.List()))
	expect := fmt.Sprintf("%v", []map[string]any{
		tuple.Merge(foo, baz),
		tuple.Merge(foo, bazpub),
		tuple.Merge(foobar, baz),
		tuple.Merge(foobar, bazpub),
	})
	if actual != expect {
		t.Errorf("expected %v got %v", expect, actual)
	}
}
