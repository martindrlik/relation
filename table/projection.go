package table

import (
	"maps"
	"slices"

	"github.com/martindrlik/rex/schema"
)

func (t *Table) Projection(p ...string) (*Table, error) {
	s, ok := t.Schema.Projection(p...)
	if !ok {
		return nil, schema.ErrMismatch
	}

	w, _ := New(s)

	for _, r := range t.RelationSet {
		i := s.Intersection(r.Schema)
		if len(i) == 0 {
			continue
		}
		s := slices.Collect(maps.Keys(i))
		wr, _ := r.Projection(s...)
		for wt := range wr.List() {
			_ = w.Add(wt)
		}
	}

	return w, nil
}
