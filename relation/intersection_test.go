package relation_test

import (
	"fmt"
	"slices"
	"testing"

	"github.com/martindrlik/rex/schema"
)

func TestIntersection(t *testing.T) {
	fb2 := map[string]any{"foo": 1, "bar": 2.0}
	fb3 := map[string]any{"foo": 1, "bar": 3.0}
	fb4 := map[string]any{"foo": 1, "bar": 4.0}
	r := newRelation(t, schema.FromTuple(fb2))
	s := newRelation(t, schema.FromTuple(fb3))
	add(t, r, fb2, fb3)
	add(t, s, fb3, fb4)
	w, err := r.Intersection(s)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	actual := fmt.Sprintf("%v", slices.Collect(w.List()))
	expect := fmt.Sprintf("%v", []map[string]any{fb3})
	if actual != expect {
		t.Errorf("expected %v got %v", expect, actual)
	}

	t.Run("mismatch", func(t *testing.T) {
		s := newRelation(t, schema.FromTuple(map[string]any{"foo": 1}))
		_, err := r.Intersection(s)
		if err != schema.ErrMismatch {
			t.Errorf("expected error %v got %v", schema.ErrMismatch, err)
		}
	})
}
