package relation_test

import (
	"fmt"
	"testing"

	"github.com/martindrlik/rex/relation"
	"github.com/martindrlik/rex/schema"
)

func TestClone(t *testing.T) {
	foobar := map[string]any{"foo": 1, "bar": 2.0}
	r := newRelation(t, schema.FromTuple(foobar), foobar)
	c := relation.Clone(r)
	if !r.Schema.Equal(c.Schema) {
		t.Errorf("expected %v got %v", r.Schema, c.Schema)
	}
	actual := fmt.Sprintf("%v", c.TupleSet)
	expect := fmt.Sprintf("%v", r.TupleSet)
	if actual != expect {
		t.Errorf("expected %v got %v", expect, actual)
	}
}
