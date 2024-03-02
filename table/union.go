package table

import (
	"github.com/martindrlik/rex/relation"
	"github.com/martindrlik/rex/require"
)

func (t *Table) Union(s *Table) (*Table, error) {
	if !t.Schema().IsEqual(s.Schema()) {
		return nil, relation.ErrSchemaMismatch
	}

	x := require.Must(New(t.Schema().Attributes()...))
	for _, r := range t.Relations() {
		s := require.Must(r.Project(t.Schema().Attributes()...))
		x.r = append(x.r, s)
	}

	for _, r := range s.Relations() {
		for _, u := range r.Tuples() {
			x.Append(u)
		}
	}
	return x, nil
}
