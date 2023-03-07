package table_test

import (
	"fmt"
	"slices"
	"testing"

	"github.com/martindrlik/rex/schema"
)

func TestRename1(t *testing.T) {
	foobaz := map[string]any{"foo": 1, "baz": 2.0}
	q := newTable(t, schema.FromTuple(foobaz), foobaz)
	r, err := q.Rename1("baz", "bar")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	foobar := map[string]any{"foo": 1, "bar": 2.0}
	actual := fmt.Sprintf("%v", slices.Collect(r.List()))
	expect := fmt.Sprintf("%v", []map[string]any{foobar})
	if actual != expect {
		t.Errorf("expected %v got %v", expect, actual)
	}

	t.Run("mismatch", func(t *testing.T) {
		_, err := q.Rename1("pub", "bar")
		if err != schema.ErrMismatch {
			t.Fatalf("unexpected error: %v", err)
		}
	})
}
