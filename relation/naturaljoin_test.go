package relation_test

import (
	"fmt"
	"slices"
	"testing"

	"github.com/martindrlik/rex/schema"
	"github.com/martindrlik/rex/tuple"
)

func TestNaturalJoin(t *testing.T) {
	t.Run("", func(t *testing.T) {
		foobar := map[string]any{"foo": 1, "bar": 2.0}
		foobaz := map[string]any{"foo": 1, "baz": "3"}
		r := newRelation(t, schema.FromTuple(foobar))
		s := newRelation(t, schema.FromTuple(foobaz))
		add(t, r, foobar)
		add(t, s, foobaz)
		w, ok := r.NaturalJoin(s)
		if !ok {
			t.Fatal("unexpected empty result")
		}
		actual := fmt.Sprintf("%v", slices.Collect(w.List()))
		expect := fmt.Sprintf("%v", []map[string]any{{"foo": 1, "bar": 2.0, "baz": "3"}})
		if actual != expect {
			t.Errorf("expected %v got %v", expect, actual)
		}
	})

	t.Run("cartasian product", func(t *testing.T) {
		foo1 := map[string]any{"foo": 1}
		foo2 := map[string]any{"foo": 2}
		bar1 := map[string]any{"bar": 1.0}
		bar2 := map[string]any{"bar": 2.0}
		r := newRelation(t, schema.FromTuple(foo1))
		s := newRelation(t, schema.FromTuple(bar1))
		add(t, r, foo1)
		add(t, r, foo2)
		add(t, s, bar1)
		add(t, s, bar2)
		w, ok := r.NaturalJoin(s)
		if !ok {
			t.Fatal("unexpected empty result")
		}
		actual := fmt.Sprintf("%v", slices.Collect(w.List()))
		expect := fmt.Sprintf("%v", []map[string]any{
			tuple.Merge(foo1, bar1),
			tuple.Merge(foo1, bar2),
			tuple.Merge(foo2, bar1),
			tuple.Merge(foo2, bar2),
		})
		if actual != expect {
			t.Errorf("expected %v got %v", expect, actual)
		}
	})

	t.Run("empty result", func(t *testing.T) {
		foo1bar := map[string]any{"foo": 1, "bar": 1.0}
		foo2bar := map[string]any{"foo": 2, "bar": 2.0}
		r := newRelation(t, schema.FromTuple(foo1bar))
		s := newRelation(t, schema.FromTuple(foo2bar))
		add(t, r, foo1bar)
		add(t, s, foo2bar)
		_, ok := r.NaturalJoin(s)
		if ok {
			t.Error("unexpected non-empty result")
		}
	})
}
