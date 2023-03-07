package schema_test

import (
	"testing"

	"github.com/martindrlik/rex/schema"
)

func TestProjection(t *testing.T) {
	s := schema.FromTuple(map[string]any{"foo": 1, "bar": 2.0})
	actual := s.Projection("foo")
	expect := schema.FromTuple(map[string]any{"foo": 1})
	if !actual.Equal(expect) {
		t.Errorf("expected %v got %v", expect, actual)
	}

	t.Run("empty", func(t *testing.T) {
		actual := s.Projection("baz")
		expect := schema.FromTuple(map[string]any{})
		if !actual.Equal(expect) {
			t.Errorf("expected empty schema as baz is not in s got %v", actual)
		}
	})
}
