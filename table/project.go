package table

import (
	"github.com/martindrlik/rex/require"
	"github.com/martindrlik/rex/tuple"
)

func (t *Table) Project(attribute string, others ...string) *Table {
	others = append([]string{attribute}, others...)
	x := require.NoError(New(others...))
	for _, r := range t.Relations(All) {
		for _, u := range r.Tuples() {
			require.NilError(x.Append(func() tuple.T {
				switch {
				case x.Schema().IsSubset(u.Schema()):
					return u.Project(x.Schema())
				case u.Schema().IsSubset(x.Schema()):
					return u.Project(u.Schema())
				default:
					return u.Project(x.Schema())
				}
			}()))
		}
	}
	return x
}
