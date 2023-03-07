package relation

import (
	"slices"

	"github.com/martindrlik/rex/schema"
)

// Union returns a new relation with tuples that are in r or in s.
func (r *Relation) Union(s *Relation) (*Relation, error) {
	if !r.Schema.Equal(s.Schema) {
		return nil, schema.ErrMismatch
	}
	w, _ := New(r.Schema)
	w.TupleSet = slices.Clone(r.TupleSet)
	for t := range s.List() {
		if !w.TupleSet.Has(t) {
			w.TupleSet.Add(t)
		}
	}
	return w, nil
}
