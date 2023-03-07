package relation_test

import (
	"testing"

	"github.com/martindrlik/rex/schema"
)

func TestDelete(t *testing.T) {
	foobar := map[string]any{"foo": 1, "bar": 2.0}
	r := newRelation(t, schema.FromTuple(foobar), foobar)
	if !r.TupleSet.Has(foobar) {
		t.Errorf("expected to contain %v", foobar)
	}
	r.Delete(foobar)
	if r.TupleSet.Has(foobar) {
		t.Errorf("expected to not contain %v", foobar)
	}
}
