package table

import (
	"github.com/martindrlik/rex/maps"
	"github.com/martindrlik/rex/relation"
	"github.com/martindrlik/rex/schema"
	"github.com/martindrlik/rex/tuple"
)

type Table struct {
	s schema.Schema
	r []*relation.Relation
}

func NewTable(attributes ...string) *Table {
	return &Table{s: schema.NewSchema(attributes...)}
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

func (t *Table) Contains(v tuple.Tuple) bool {
	s := schema.NewSchema(maps.Keys(v)...)
	r := t.relationBySchema(s)
	if r == nil {
		return false
	}
	return r.Contains(v)
}

func (t *Table) Append(u tuple.Tuple) error {
	s := schema.NewSchema(maps.Keys(u)...)
	if !t.Schema().IsEqual(s) && !t.Schema().IsSubset(s) {
		return relation.ErrSchemaMismatch
	}
	r := t.relationBySchema(s)
	if r != nil {
		return r.Append(u)
	}
	r, err := relation.NewRelation(u)
	if err != nil {
		return err
	}
	t.r = append(t.r, r)
	return nil
}
