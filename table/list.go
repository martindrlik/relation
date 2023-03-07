package table

import "iter"

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
