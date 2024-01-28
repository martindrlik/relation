package rex

import (
	"github.com/martindrlik/rex/schema"
)

type Table struct {
	schema    map[string]struct{}
	relations []*R
}

func NewTable(schema ...string) *Table { return newTable(schema) }
func newTable(s []string) *Table       { return &Table{schema: schema.Map(s...)} }

func (tbl *Table) Schema() map[string]struct{}  { return tbl.schema }
func (tbl *Table) Relations() []*R              { return tbl.relations }
func (tbl *Table) Pick(schema ...string) *Table { return tbl.pick(schema) }

func (tbl *Table) pick(s []string) *Table {
	if len(s) == 0 {
		return tbl
	}
	m := schema.Map(s...)
	rs := []*R{}
	for _, r := range tbl.relations {
		if schema.IsEqual(r.Schema(), m) || schema.IsSubset(r.Schema(), m) {
			rs = append(rs, r)
		}
	}
	p := newTable(schema.Slice(relationsSchema(rs)))
	p.relations = rs
	return p
}

func relationsSchema(rs []*R) map[string]struct{} {
	m := map[string]struct{}{}
	for _, r := range rs {
		for k := range r.Schema() {
			m[k] = struct{}{}
		}
	}
	return m
}

func (tbl *Table) Equal(other *Table) bool {
	if len(tbl.relations) != len(other.relations) ||
		!schema.IsEqual(tbl.schema, other.schema) {
		return false
	}
	for i, r := range tbl.relations {
		if !r.Equal(other.relations[i]) {
			return false
		}
	}
	return true
}

func (tbl *Table) Add(t T) *Table {
	isCompatible := schema.IsEqual(t, tbl.schema) || schema.IsSubset(t, tbl.schema)
	if !isCompatible {
		panic("schema mismatch")
	}

	r := tbl.tryFindCompatible(t)
	if r == nil {
		r = &R{}
		tbl.relations = append(tbl.relations, r)
	}
	r.Add(t)
	return tbl
}

func (tbl *Table) tryFindCompatible(t T) *R {
	for _, r := range tbl.relations {
		if schema.IsEqual(t, (*r)[0]) {
			return r
		}
	}
	return nil
}

func Union(t1, t2 *Table) *Table {
	if !schema.IsEqual(t1.schema, t2.schema) {
		return &Table{}
	}
	tbl := NewTable(schema.Slice(t1.Schema())...)
	add := func(t T) { tbl.Add(t) }
	t1.forEach(add)
	t2.forEach(add)
	return tbl
}

func (tbl *Table) forEach(f func(t T)) {
	for _, r := range tbl.relations {
		for _, t := range *r {
			f(t)
		}
	}
}
