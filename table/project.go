package table

import (
	"github.com/martindrlik/rex/relation"
	"github.com/martindrlik/rex/require"
)

func (t *Table) Project(attributes ...string) (*Table, error) {
	if len(attributes) == 0 {
		return nil, relation.ErrMissingSchema
	}

	x := require.Must(NewTable(attributes...))
	for _, r := range t.Relations() {
		s, err := r.Project(attributes...)
		if err != nil {
			return nil, err
		}
		for _, u := range s.Tuples() {
			err := x.Append(u)
			if err != nil {
				return nil, err
			}
		}
	}
	return x, nil
}
