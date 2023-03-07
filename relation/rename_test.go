package relation_test

import (
	"fmt"
	"slices"
	"testing"

	"github.com/martindrlik/rex/schema"
)

func TestRename1(t *testing.T) {
	foobaz := map[string]any{"foo": 1, "baz": 2.0}
	foobar := map[string]any{"foo": 1, "bar": 2.0}
	r := newRelation(t, schema.FromTuple(foobaz), foobaz)
	w, err := r.Rename1("baz", "bar")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	actual := fmt.Sprintf("%v", slices.Collect(w.List()))
	expect := fmt.Sprintf("%v", []map[string]any{foobar})
	if actual != expect {
		t.Errorf("expected %v got %v", expect, actual)
	}

	t.Run("mismatch", func(t *testing.T) {
		_, err := r.Rename1("pub", "bar")
		if err != schema.ErrMismatch {
			t.Errorf("expected error %v got %v", schema.ErrMismatch, err)
		}
		_, err = r.Rename1("baz", "foo")
		if err != schema.ErrMismatch {
			t.Errorf("expected error %v got %v", schema.ErrMismatch, err)
		}
	})
}
