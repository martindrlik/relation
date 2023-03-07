package table

import (
	"github.com/martindrlik/rex/relation"
	"github.com/martindrlik/rex/schema"
)

type Table struct {
	schema.Schema
	relation.RelationSet
}

// New returns a new table with the given schema.
func New(s schema.Schema) (*Table, error) {
	if len(s) == 0 {
		return nil, schema.ErrEmpty
	}
	return &Table{Schema: s}, nil
}

// Add adds a tuple to the table. If the tuple is already in the table, it does nothing.
func (t *Table) Add(u map[string]any) error {
	if len(u) == 0 {
		return schema.ErrEmpty
	}
	us := schema.FromTuple(u)
	if !us.IsSubsetOf(t.Schema) {
		return schema.ErrMismatch
	}
	r, ok := t.Relation(us)
	if !ok {
		r, _ = relation.New(us)
		t.RelationSet.Add(r)
	}
	if !r.TupleSet.Has(u) {
		r.TupleSet.Add(u)
	}
	return nil
}

func (t *Table) Has(u map[string]any) bool {
	r, ok := t.Relation(schema.FromTuple(u))
	return ok && r.TupleSet.Has(u)
}
