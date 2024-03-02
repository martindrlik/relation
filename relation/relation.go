package relation

import (
	"github.com/martindrlik/rex/maps"
	"github.com/martindrlik/rex/schema"
	"github.com/martindrlik/rex/tuple"
)

type Relation struct {
	s schema.Schema
	t []tuple.T
}

func New(t tuple.T, others ...tuple.T) (*Relation, error) {
	if len(t) == 0 {
		return nil, ErrMissingSchema
	}

	x := &Relation{
		s: schema.New(maps.Keys(t)...),
		t: []tuple.T{t},
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
func (r *Relation) Tuples() []tuple.T     { return r.t }

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

func (r *Relation) Contains(u tuple.T) bool {
	for _, v := range r.t {
		if u.Equals(v) {
			return true
		}
	}
	return false
}

func (r *Relation) Append(u tuple.T, tuples ...tuple.T) error {
	tuples = append([]tuple.T{u}, tuples...)
	for _, u := range tuples {
		if !r.Schema().IsEqual(u.Schema()) {
			return ErrSchemaMismatch
		}
	}
	for _, u := range tuples {
		r.append(u)
	}
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

func (r *Relation) append(t tuple.T) {
	if !r.Contains(t) {
		r.t = append(r.t, t)
	}
}
