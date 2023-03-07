package schema_test

import (
	"maps"
	"reflect"
	"testing"

	"github.com/martindrlik/rex/schema"
)

func TestFromTuple(t *testing.T) {
	actual := schema.FromTuple(map[string]any{"foo": 1, "bar": 2.0})
	expect := map[string]reflect.Type{"foo": reflect.TypeOf(1), "bar": reflect.TypeOf(2.0)}

	if !maps.Equal(actual, expect) {
		t.Errorf("expected %v got %v", expect, actual)
	}
}

func TestEqual(t *testing.T) {
	schema1 := schema.FromTuple(map[string]any{"foo": 1, "bar": 2.0})
	schema2 := schema.FromTuple(map[string]any{"foo": 1, "bar": 2.0})

	if !schema1.Equal(schema2) {
		t.Errorf("expected %v and %v to contain the same key/value pairs", schema1, schema2)
	}
}

func TestIntersection(t *testing.T) {
	schema1 := schema.FromTuple(map[string]any{"foo": 1, "bar": 2.0})
	schema2 := schema.FromTuple(map[string]any{"bar": 2.0, "baz": "3"})

	actual := schema1.Intersection(schema2)
	expect := schema.FromTuple(map[string]any{"bar": 2.0})

	if !maps.Equal(actual, expect) {
		t.Errorf("expected %v got %v", expect, actual)
	}
}

func TestHas(t *testing.T) {
	s := schema.FromTuple(map[string]any{"foo": 1, "bar": 2.0})

	if !s.Has("foo") {
		t.Error("expected schema to contain key foo")
	}

	if s.Has("baz") {
		t.Error("expected schema to not contain key baz")
	}
}
