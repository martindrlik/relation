package table_test

import (
	"testing"

	"github.com/martindrlik/rex/schema"
)

func TestDelete(t *testing.T) {
	foobar := map[string]any{"foo": 1, "bar": 2.0}
	q := newTable(t, schema.FromTuple(foobar), foobar)
	if !q.Has(foobar) {
		t.Errorf("expected to contain %v", foobar)
	}
	q.Delete(foobar)
	if q.Has(foobar) {
		t.Errorf("expected to not contain %v", foobar)
	}
}
