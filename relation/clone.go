package relation

import "slices"

// Clone returns a new relation with the same schema and tuples as r.
// The schema is considered read-only so it is not cloned.
func Clone(r *Relation) *Relation {
	w, _ := New(r.Schema)
	w.TupleSet = slices.Clone(r.TupleSet)
	return w
}
