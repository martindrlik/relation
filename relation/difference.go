package relation

import "github.com/martindrlik/rex/schema"

// Difference returns a new relation with tuples that are in r but not in s.
func (r *Relation) Difference(s *Relation) (*Relation, error) {
	if !r.Schema.Equal(s.Schema) {
		return nil, schema.ErrMismatch
	}
	w, _ := New(r.Schema)
	for t := range r.List() {
		if !s.TupleSet.Has(t) {
			w.TupleSet.Add(t)
		}
	}
	return w, nil
}
