package relation

import (
	"github.com/martindrlik/rex/maps"
	"github.com/martindrlik/rex/schema"
	"github.com/martindrlik/rex/tuple"
)

type Relation struct {
	s schema.Schema
	t []tuple.Tuple
}

func NewRelation(t tuple.Tuple, ts ...tuple.Tuple) (*Relation, error) {
	r := &Relation{
		s: schema.NewSchema(maps.Keys(t)...),
		t: []tuple.Tuple{t},
	}
	for _, t := range ts {
		if err := r.Append(t); err != nil {
			return nil, err
		}
	}
	return r, nil
}

func (r *Relation) Schema() schema.Schema { return r.s }
func (r *Relation) Tuples() []tuple.Tuple { return r.t }

func (r *Relation) Equals(other *Relation) bool {
	if !r.s.IsEqual(other.s) {
		return false
	}
	if len(r.t) != len(other.t) {
		return false
	}
	for _, t := range r.t {
		if !other.Contains(t) {
			return false
		}
	}
	return true
}

func (r *Relation) Contains(v tuple.Tuple) bool {
	for _, u := range r.t {
		if u.Equals(v) {
			return true
		}
	}
	return false
}

func (r *Relation) Append(t tuple.Tuple) error {
	if !r.s.IsEqual(schema.NewSchema(maps.Keys(t)...)) {
		return ErrSchemaMismatch
	}
	if !r.Contains(t) {
		r.t = append(r.t, t)
	}
	return nil
}
