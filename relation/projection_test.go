package relation_test

import (
	"fmt"
	"slices"
	"testing"

	"github.com/martindrlik/rex/schema"
)

func TestProjection(t *testing.T) {
	fbb := map[string]any{"foo": 1, "bar": 2.0, "baz": "3"}
	r := newRelation(t, schema.FromTuple(fbb))
	add(t, r, fbb)

	s, err := r.Projection("foo", "baz")
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}
	actual := fmt.Sprintf("%v", slices.Collect(s.List()))
	expect := "[map[baz:3 foo:1]]"
	if actual != expect {
		t.Errorf("expected %v got %v", expect, actual)
	}

	t.Run("mismatch", func(t *testing.T) {
		_, err := r.Projection("foo", "qux")
		if err != schema.ErrMismatch {
			t.Errorf("expected error %v got %v", schema.ErrMismatch, err)
		}
	})
}
