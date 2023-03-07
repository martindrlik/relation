package relation_test

import (
	"testing"

	"github.com/martindrlik/rex/relation"
	"github.com/martindrlik/rex/schema"
)

func TestRelationSet(t *testing.T) {
	rs := relation.RelationSet{}
	r1 := newRelation(t, schema.FromTuple(map[string]any{"foo": 1}))
	rs.Add(r1)
	r, ok := rs.Relation(r1.Schema)
	if !ok {
		t.Error("expected to find relation in set")
	}
	if r != r1 {
		t.Error("expected to find the same relation in set")
	}
	t.Run("not found", func(t *testing.T) {
		_, ok := rs.Relation(schema.FromTuple(map[string]any{"bar": 2}))
		if ok {
			t.Error("expected not to find relation in set")
		}
	})
}
