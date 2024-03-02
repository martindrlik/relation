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

func NewRelation(t tuple.Tuple, others ...tuple.Tuple) (*Relation, error) {
	if len(t) == 0 {
		return nil, ErrMissingSchema
	}

	x := &Relation{
		s: schema.NewSchema(maps.Keys(t)...),
		t: []tuple.Tuple{t},
	}
	for _, o := range others {
		if !x.Schema().IsEqual(o.Schema()) {
			return nil, ErrSchemaMismatch
		}
		x.append(o)
	}
	return x, nil
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

func (r *Relation) Contains(u tuple.Tuple) bool {
	for _, v := range r.t {
		if u.Equals(v) {
			return true
		}
	}
	return false
}

func (r *Relation) Append(u tuple.Tuple) error {
	if !r.Schema().IsEqual(u.Schema()) {
		return ErrSchemaMismatch
	}
	r.append(u)
	return nil
}

func (r *Relation) appendRelation(s *Relation) error {
	if !r.Schema().IsEqual(s.Schema()) {
		return ErrSchemaMismatch
	}
	for _, t := range s.t {
		r.append(t)
	}
	return nil
}

func (r *Relation) append(t tuple.Tuple) {
	if !r.Contains(t) {
		r.t = append(r.t, t)
	}
}
