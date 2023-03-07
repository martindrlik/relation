package table

import (
	"github.com/martindrlik/rex/relation"
	"github.com/martindrlik/rex/schema"
)

func (t *Table) Difference(u *Table) (*Table, error) {
	if !t.Schema.Equal(u.Schema) {
		return nil, schema.ErrMismatch
	}

	w, _ := New(t.Schema)
	for _, tr := range t.RelationSet {
		if tr.Schema.Equal(t.Schema) {
			if ur, ok := u.Relation(tr.Schema); ok {
				if w0, _ := tr.Difference(ur); len(w0.TupleSet) > 0 {
					w.RelationSet.Add(w0)
				}
			}
			continue
		}
		if len(tr.TupleSet) > 0 {
			w.RelationSet.Add(relation.Clone(tr))
		}
	}

	return w, nil
}
