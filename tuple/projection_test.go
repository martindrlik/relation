package tuple_test

import (
	"maps"
	"testing"

	"github.com/martindrlik/rex/tuple"
)

func TestProjection(t *testing.T) {
	u := tuple.Tuple{"foo": 1, "bar": 2.0, "baz": "3"}
	actual := u.Projection("foo", "baz")
	expect := tuple.Tuple{"foo": 1, "baz": "3"}
	if !maps.Equal(actual, expect) {
		t.Errorf("expected %v got %v", expect, actual)
	}
}
