package table

import (
	"github.com/martindrlik/rex/relation"
	"github.com/martindrlik/rex/schema"
)

// Union returns a new table that is the union of t and u.
func (t *Table) Union(u *Table) (*Table, error) {
	if !t.Schema.Equal(u.Schema) {
		return nil, schema.ErrMismatch
	}

	w, _ := New(t.Schema)
	for _, tr := range t.RelationSet {
		ur, ok := u.Relation(tr.Schema)
		if ok {
			wr, _ := tr.Union(ur)
			w.RelationSet.Add(wr)
		} else {
			w.RelationSet.Add(relation.Clone(tr))
		}
	}
	for _, ur := range u.RelationSet {
		_, ok := t.Relation(ur.Schema)
		if !ok {
			w.RelationSet.Add(relation.Clone(ur))
		}
	}

	return w, nil
}
