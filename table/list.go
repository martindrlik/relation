package table

import "iter"

// List returns a sequence of all tuples in the table.
func (t *Table) List() iter.Seq[map[string]any] {
	return func(yield func(map[string]any) bool) {
		for _, r := range t.RelationSet {
			for _, t := range r.TupleSet {
				if !yield(t) {
					return
				}
			}
		}
	}
}
