package relation

import (
	"github.com/martindrlik/rex/schema"
	"github.com/martindrlik/rex/tuple"
)

func (r *Relation) Projection(p ...string) (*Relation, error) {
	s, ok := r.Schema.Projection(p...)
	if !ok {
		return nil, schema.ErrMismatch
	}
	w, _ := New(s)
	for rt := range r.List() {
		wt := tuple.Tuple(rt).Projection(p...)
		if !w.TupleSet.Has(wt) {
			w.TupleSet.Add(wt)
		}
	}
	return w, nil
}
