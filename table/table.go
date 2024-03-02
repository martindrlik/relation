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

func (t *Table) Schema() schema.Schema           { return t.s }
func (t *Table) Relations() []*relation.Relation { return t.r }

func (t *Table) Equals(other *Table) bool {
	if !t.s.IsEqual(other.s) {
		return false
	}
	if len(t.r) != len(other.r) {
		return false
	}
	for _, r := range t.r {
		s := other.relationBySchema(r.Schema())
		if s == nil {
			return false
		}
		if !r.Equals(s) {
			return false
		}
	}
	return true
}

func (t *Table) relationBySchema(s schema.Schema) *relation.Relation {
	for _, r := range t.r {
		if r.Schema().IsEqual(s) {
			return r
		}
	}
	return nil
}

func (t *Table) Contains(u tuple.T) bool {
	s := schema.New(maps.Keys(u)...)
	r := t.relationBySchema(s)
	if r == nil {
		return false
	}
	return r.Contains(u)
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
		r := t.relationBySchema(u.Schema())
		if r != nil {
			return r.Append(u)
		}
		r = require.Must(relation.New(u))
		t.r = append(t.r, r)
	}
	return nil
}
