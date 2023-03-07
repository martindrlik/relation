package relation

import "iter"

// List returns a sequence of tuples in the relation.
func (r *Relation) List() iter.Seq[map[string]any] {
	return func(yield func(map[string]any) bool) {
		for _, t := range r.TupleSet {
			if !yield(t) {
				return
			}
		}
	}
}
