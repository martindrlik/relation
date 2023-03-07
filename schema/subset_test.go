package schema_test

import (
	"testing"

	"github.com/martindrlik/rex/schema"
)

func TestIsSubsetOf(t *testing.T) {
	u := schema.FromTuple(map[string]any{"foo": 1, "bar": 2.0})
	v := schema.FromTuple(map[string]any{"foo": 1})
	if !v.IsSubsetOf(u) {
		t.Errorf("expected %v to be a subset of %v", v, u)
	}

	t.Run("equal", func(t *testing.T) {
		if !u.IsSubsetOf(u) {
			t.Errorf("expected %v to be a subset of %v", u, u)
		}
	})

	t.Run("not a subset", func(t *testing.T) {
		v := schema.FromTuple(map[string]any{"foo": 1, "baz": "3"})
		if u.IsSubsetOf(v) {
			t.Errorf("expected %v not to be a subset of %v", u, v)
		}
	})

	t.Run("no a subset 2", func(t *testing.T) {
		v := schema.FromTuple(map[string]any{"foo": 1, "bar": 2.0, "baz": "3"})
		if v.IsSubsetOf(u) {
			t.Errorf("expected %v not to be a subset of %v", v, u)
		}
	})

	t.Run("empty", func(t *testing.T) {
		if !schema.FromTuple(map[string]any{}).IsSubsetOf(u) {
			t.Errorf("expected empty schema to be a subset of %v", u)
		}
	})
}
