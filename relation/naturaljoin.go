package relation

import (
	"maps"

	"github.com/martindrlik/rex/tuple"
)

func (r *Relation) NaturalJoin(s *Relation) (*Relation, bool) {
	common := r.Schema.Intersection(s.Schema)
	concat := func(u, v map[string]any) (map[string]any, bool) {
		w := make(map[string]any)
		for a := range common {
			if u[a] != v[a] {
				return nil, false
			}
		}
		maps.Copy(w, u)
		maps.Copy(w, v)
		return w, true
	}

	w := func() *Relation {
		cs := tuple.Merge(r.Schema, s.Schema)
		w, _ := New(cs)
		return w
	}()
	for rt := range r.List() {
		for st := range s.List() {
			if t, ok := concat(rt, st); ok {
				w.TupleSet.Add(t)
			}
		}
	}
	return w, len(w.TupleSet) > 0
}
