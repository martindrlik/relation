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

	s, ok := r.Projection("foo", "baz")
	if !ok {
		t.Fatal("unexpected empty result")
	}
	actual := fmt.Sprintf("%v", slices.Collect(s.List()))
	expect := "[map[baz:3 foo:1]]"
	if actual != expect {
		t.Errorf("expected %v got %v", expect, actual)
	}

	t.Run("empty result", func(t *testing.T) {
		_, ok := r.Projection("foo", "qux")
		if ok {
			t.Error("unexpected non-empty result")
		}
	})
}
