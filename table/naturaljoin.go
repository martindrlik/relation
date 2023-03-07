package table

import (
	"github.com/martindrlik/rex/relation"
	"github.com/martindrlik/rex/schema"
	"github.com/martindrlik/rex/tuple"
)

func (t *Table) NaturalJoin(u *Table) *Table {
	common := t.Schema.Intersection(u.Schema)
	if len(common) == 0 {
		return t.cartasianProduct(u)
	} else {
		return t.naturalJoin(u, common)
	}
}

func (t *Table) cartasianProduct(u *Table) *Table {
	w, _ := New(tuple.Merge(t.Schema, u.Schema))
	for _, tr := range t.RelationSet {
		for _, ur := range u.RelationSet {
			if wr, ok := tr.NaturalJoin(ur); ok {
				w.RelationSet.Add(wr)
			}
		}
	}
	return w
}

func (t *Table) naturalJoin(u *Table, common schema.Schema) *Table {
	rr := []*relation.Relation{}
	for _, tr := range t.RelationSet {
		for _, ur := range u.RelationSet {
			if common.IsSubsetOf(tr.Schema) && common.IsSubsetOf(ur.Schema) {
				if wr, ok := tr.NaturalJoin(ur); ok {
					rr = append(rr, wr)
				}
			}
		}
	}

	w, _ := New(tuple.Merge(t.Schema, u.Schema))
	for _, r := range rr {
		w.RelationSet.Add(r)
	}
	return w
}
