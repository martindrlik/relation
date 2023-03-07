package relation_test

import (
	"fmt"
	"slices"
	"testing"

	"github.com/martindrlik/rex/schema"
)

func TestUnion(t *testing.T) {
	foo1 := map[string]any{"foo": 1}
	foo2 := map[string]any{"foo": 2}
	r := newRelation(t, schema.FromTuple(foo1))
	s := newRelation(t, schema.FromTuple(foo2))
	add(t, r, foo1)
	add(t, s, foo2)
	w, err := r.Union(s)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	actual := fmt.Sprintf("%v", slices.Collect(w.List()))
	expect := fmt.Sprintf("%v", []map[string]any{foo1, foo2})

	if actual != expect {
		t.Errorf("expected %v got %v", expect, actual)
	}

	t.Run("mismatch", func(t *testing.T) {
		bar := map[string]any{"bar": 2.0}
		s := newRelation(t, schema.FromTuple(bar))
		_, err := r.Union(s)
		if err != schema.ErrMismatch {
			t.Errorf("expected error %v got %v", schema.ErrMismatch, err)
		}
	})
}
