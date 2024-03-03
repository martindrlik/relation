package table

import (
	"github.com/martindrlik/rex/relation"
	"github.com/martindrlik/rex/schema"
)

func All[T any](T) bool {
	return true
}

func Matching(schema schema.Schema) func(*relation.Relation) bool {
	return func(r *relation.Relation) bool {
		return r.Schema().IsEqual(schema)
	}
}

func Missing(attributes ...string) func(*relation.Relation) bool {
	schema := schema.New(attributes...)
	return func(r *relation.Relation) bool {
		return !schema.IsSubset(r.Schema())
	}
}

func (t *Table) Relations(predicate func(*relation.Relation) bool) []*relation.Relation {
	x := []*relation.Relation{}
	for _, r := range t.r {
		if predicate(r) {
			x = append(x, r)
		}
	}
	return x
}
