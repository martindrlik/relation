package schema_test

import (
	"testing"

	"github.com/martindrlik/rex/schema"
)

func TestProjection(t *testing.T) {
	s := schema.FromTuple(map[string]any{"foo": 1, "bar": 2.0})
	actual, ok := s.Projection("foo")
	if !ok {
		t.Fatal("unexpected empty result")
	}
	expect := schema.FromTuple(map[string]any{"foo": 1})
	if !actual.Equal(expect) {
		t.Errorf("expected %v got %v", expect, actual)
	}

	t.Run("empty", func(t *testing.T) {
		_, ok := s.Projection("baz")
		if ok {
			t.Error("unexpected non-empty result")
		}
	})
}
