package rex

import (
	"sort"

	"github.com/martindrlik/rex/schema"
)

type Table struct {
	schema    map[string]any
	relations []*R
}

func NewTable(schema ...string) *Table {
	return &Table{schema: schemaMap(schema)}
}

func schemaMap(schema []string) map[string]any {
	m := make(map[string]any)
	for _, a := range schema {
		m[a] = struct{}{}
	}
	return m
}

func (tbl *Table) Schema() []string {
	s := make([]string, 0, len(tbl.schema))
	for k := range tbl.schema {
		s = append(s, k)
	}
	sort.Strings(s)
	return s
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
	tbl := NewTable(t1.Schema()...)
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
