package relation

import (
	"github.com/martindrlik/rex/tuple"
)

func (r *Relation) Projection(p ...string) (*Relation, bool) {
	s := r.Schema.Projection(p...)
	if len(s) != len(p) || !s.IsSubsetOf(r.Schema) {
		return nil, false
	}
	w, _ := New(s)
	for rt := range r.List() {
		wt := tuple.Tuple(rt).Projection(p...)
		if !w.TupleSet.Has(wt) {
			w.TupleSet.Add(wt)
		}
	}
	return w, true
}
