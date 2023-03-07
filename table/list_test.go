package table_test

import (
	"fmt"
	"slices"
	"testing"

	"github.com/martindrlik/rex/schema"
	"github.com/martindrlik/rex/tuple"
)

func TestList(t *testing.T) {
	foo := map[string]any{"foo": 1}
	bar := map[string]any{"bar": 1.0}
	foobar := tuple.Merge(foo, bar)
	q := newTable(t, schema.FromTuple(foobar), foo, bar, foobar)
	actual := fmt.Sprintf("%v", slices.Collect(q.List()))
	expect := fmt.Sprintf("%v", []map[string]any{foo, bar, foobar})
	if actual != expect {
		t.Errorf("expected %v got %v", expect, actual)
	}
	t.Run("break", func(t *testing.T) {
		uu := []map[string]any{}
		for u := range q.List() {
			uu = append(uu, u)
			if len(uu) == 2 {
				break
			}
		}
		actual := fmt.Sprintf("%v", uu)
		expect := fmt.Sprintf("%v", []map[string]any{foo, bar})
		if actual != expect {
			t.Errorf("expected %v got %v", expect, actual)
		}
	})
}
