package table

import (
	"github.com/martindrlik/rex/relation"
	"github.com/martindrlik/rex/schema"
	"github.com/martindrlik/rex/tuple"
)

func (t *Table) Projection(p ...string) (*Table, bool) {
	ts := tuple.TupleSet{}
	for _, r := range t.RelationSet {
		w, ok := r.Projection(p...)
		if !ok {
			continue
		}
		for wt := range w.List() {
			if !ts.Has(wt) {
				ts.Add(wt)
			}
		}
	}

	if len(ts) == 0 {
		return nil, false
	}

	wr, _ := relation.New(schema.FromTuple(ts[0]))
	wr.TupleSet = ts
	w, _ := New(wr.Schema)
	w.RelationSet.Add(wr)
	return w, true
}
