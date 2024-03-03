package table

import (
	"github.com/martindrlik/rex/maps"
	"github.com/martindrlik/rex/relation"
	"github.com/martindrlik/rex/require"
	"github.com/martindrlik/rex/schema"
	"github.com/martindrlik/rex/tuple"
)

type Table struct {
	s schema.Schema
	r []*relation.Relation
}

func New(attributes ...string) (*Table, error) {
	if len(attributes) == 0 {
		return nil, relation.ErrMissingSchema
	}
	x := &Table{s: schema.New(attributes...)}
	return x, nil
}

func (t *Table) Schema() schema.Schema { return t.s }

func (t *Table) Equals(other *Table) bool {
	if !t.s.IsEqual(other.s) {
		return false
	}
	if len(t.r) != len(other.r) {
		return false
	}
	for _, r := range t.r {
		rs := other.Relations(Matching(r.Schema()))
		if len(rs) != 1 {
			return false
		}
		if !r.Equals(rs[0]) {
			return false
		}
	}
	return true
}

func (t *Table) Contains(u tuple.T) bool {
	s := schema.New(maps.Keys(u)...)
	rs := t.Relations(Matching(s))
	return len(rs) == 1 && rs[0].Contains(u)
}

func (t *Table) Append(u tuple.T, tuples ...tuple.T) error {
	tuples = append([]tuple.T{u}, tuples...)
	for _, u := range tuples {
		if !(u.Schema().IsEqual(t.Schema()) ||
			u.Schema().IsSubset(t.Schema())) {
			return relation.ErrSchemaMismatch
		}
	}
	for _, u := range tuples {
		rs := t.Relations(Matching(u.Schema()))

		if len(rs) != 1 {
			r := require.NoError(relation.New(u))
			t.r = append(t.r, r)
			continue
		}

		if err := rs[0].Append(u); err != nil {
			return err
		}
	}
	return nil
}
