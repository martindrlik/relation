package relation_test

import (
	"fmt"
	"testing"

	"github.com/martindrlik/rex/relation"
	"github.com/martindrlik/rex/schema"
)

func TestRelation(t *testing.T) {
	foobar := map[string]any{"foo": 1, "bar": 2.0}
	t.Run("schema", func(t *testing.T) {
		r := newRelation(t, schema.FromTuple(foobar))
		if !r.Schema.Equal(schema.FromTuple(foobar)) {
			t.Error("relation is created without expected schema")
		}
	})
	t.Run("add", func(t *testing.T) {
		r := newRelation(t, schema.FromTuple(foobar))
		err := r.Add(foobar)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if a, e := fmt.Sprintf("%v", *r), "{map[bar:float64 foo:int] [map[bar:2 foo:1]]}"; a != e {
			t.Errorf("expected %v got %v", e, a)
		}
	})
	t.Run("duplicate", func(t *testing.T) {
		r := newRelation(t, schema.FromTuple(foobar))
		err := r.Add(foobar)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		err = r.Add(foobar)
		if err != nil {
			t.Errorf("expected no error got %v", err)
		}
	})
	t.Run("mismatch", func(t *testing.T) {
		r := newRelation(t, schema.FromTuple(foobar))
		err := r.Add(map[string]any{"foo": 1})
		if err != schema.ErrMismatch {
			t.Fatalf("expected error %v got %v", schema.ErrMismatch, err)
		}
	})
}

func newRelation(t *testing.T, s schema.Schema, tt ...map[string]any) *relation.Relation {
	t.Helper()
	r, err := relation.New(s)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	add(t, r, tt...)
	return r
}

func TestNew(t *testing.T) {
	_, err := relation.New(nil)
	if err != schema.ErrEmpty {
		t.Errorf("expected error %v got %v", schema.ErrEmpty, err)
	}
}

func add(t *testing.T, r *relation.Relation, tt ...map[string]any) {
	t.Helper()
	for _, tup := range tt {
		err := r.Add(tup)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	}
}
