package relation_test

import (
	"fmt"
	"slices"
	"testing"

	"github.com/martindrlik/rex/schema"
)

func TestList(t *testing.T) {
	foo1 := map[string]any{"foo": 1}
	foo2 := map[string]any{"foo": 2}
	r := newRelation(t, schema.FromTuple(foo1))
	add(t, r, foo1)
	add(t, r, foo2)
	actual := fmt.Sprintf("%v", slices.Collect(r.List()))
	expect := fmt.Sprintf("%v", []map[string]any{foo1, foo2})
	if actual != expect {
		t.Errorf("expected %v got %v", expect, actual)
	}
}

func TestListPartial(t *testing.T) {
	foo1 := map[string]any{"foo": 1}
	foo2 := map[string]any{"foo": 2}
	r := newRelation(t, schema.FromTuple(foo1))
	add(t, r, foo1)
	add(t, r, foo2)
	s := make([]map[string]any, 0)
	for t := range r.List() {
		s = append(s, t)
		break
	}
	actual := fmt.Sprintf("%v", s)
	expect := fmt.Sprintf("%v", []map[string]any{foo1})
	if actual != expect {
		t.Errorf("expected %v got %v", expect, actual)
	}
}
