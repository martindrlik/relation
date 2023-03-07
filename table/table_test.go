package table_test

import (
	"testing"

	"github.com/martindrlik/rex/schema"
	"github.com/martindrlik/rex/table"
	"github.com/martindrlik/rex/tuple"
)

func TestTable(t *testing.T) {
	t.Run("empty schema", func(t *testing.T) {
		_, err := table.New(schema.Schema{})
		if err != schema.ErrEmpty {
			t.Errorf("expected error %v got %v", schema.ErrEmpty, err)
		}
		q, err := table.New(schema.FromTuple(map[string]any{"foo": 1}))
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if err := q.Add(map[string]any{}); err != schema.ErrEmpty {
			t.Errorf("expected error %v got %v", schema.ErrEmpty, err)
		}
	})
	t.Run("mismatch", func(t *testing.T) {
		q, err := table.New(schema.FromTuple(map[string]any{"foo": 1}))
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if err := q.Add(map[string]any{"bar": 2.0}); err != schema.ErrMismatch {
			t.Errorf("expected error %v got %v", schema.ErrMismatch, err)
		}
	})
	foo := map[string]any{"foo": 1}
	bar := map[string]any{"bar": 2.0}
	foobar := tuple.Merge(foo, bar)
	q, err := table.New(schema.FromTuple(foobar))
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	add := func(u map[string]any) {
		t.Helper()
		if err := q.Add(u); err != nil {
			t.Errorf("unexpected error: %v", err)
		}
	}
	add(foobar)
	add(foo)
	add(bar)

	has := func(u map[string]any) {
		t.Helper()
		r, ok := q.Relation(schema.FromTuple(u))
		if !ok {
			t.Error("expected to find relation in set")
		}
		if !r.TupleSet.Has(u) {
			t.Errorf("expected to find %v in relation", u)
		}
	}
	has(foobar)
	has(foo)
	has(bar)
}

func TestHas(t *testing.T) {
	foobar := map[string]any{"foo": 1, "bar": 2.0}
	q := newTable(t, schema.FromTuple(foobar), foobar)
	if !q.Has(foobar) {
		t.Error("expected to find tuple in table")
	}

	foo := map[string]any{"foo": 1}
	if q.Has(foo) {
		t.Errorf("expected to not find %v in table", foo)
	}

	fbb := map[string]any{"foo": 1, "bar": 2.0, "baz": "3"}
	if q.Has(fbb) {
		t.Errorf("expected to not find %v in table", fbb)
	}
}

func newTable(t *testing.T, s schema.Schema, tt ...map[string]any) *table.Table {
	t.Helper()
	q, err := table.New(s)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	add(t, q, tt...)
	return q
}

func add(t *testing.T, q *table.Table, tt ...map[string]any) {
	t.Helper()
	for _, u := range tt {
		if err := q.Add(u); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	}
}
