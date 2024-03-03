package relation

import (
	"github.com/martindrlik/rex/require"
	"github.com/martindrlik/rex/tuple"
)

func (r *Relation) NaturalJoin(other *Relation) (*Relation, error) {
	x := []tuple.T{}
	ef := equalFunc(r, other)
	for _, u := range r.Tuples() {
		for _, v := range other.Tuples() {
			if ef(u, v) {
				x = append(x, u.Join(v))
			}
		}
	}
	if len(x) == 0 {
		return nil, ErrResultIsEmpty
	}
	return require.NoError(New(x[0], x[1:]...)), nil
}

func equalFunc(r, s *Relation) func(tuple.T, tuple.T) bool {
	is := r.Schema().Intersection(s.Schema()).Attributes()
	return func(u, v tuple.T) bool {
		if len(is) == 0 {
			return true
		}
		for _, i := range is {
			if u[i] != v[i] {
				return false
			}
		}
		return true
	}
}
